package web3

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/event"
	"github.com/rs/zerolog/log"

	"github.com/CoopHive/hive/pkg/system"
	"github.com/CoopHive/hive/pkg/web3/bindings/payments"
)

type PaymentEventChannels struct {
	paymentChan chan *payments.PaymentsPayment
	paymentSubs []func(payments.PaymentsPayment)
}

func NewPaymentEventChannels() *PaymentEventChannels {
	return &PaymentEventChannels{
		paymentChan: make(chan *payments.PaymentsPayment),
		paymentSubs: []func(payments.PaymentsPayment){},
	}
}

func (p *PaymentEventChannels) Start(
	sdk *Web3SDK,
	ctx context.Context,
	cm *system.CleanupManager,
) (err error) {
	defer func() {
		eventErrorHandler(err)
	}()

	blockNumber, err := sdk.getBlockNumber()
	if err != nil {
		return err
	}

	var paymentSub event.Subscription

	connectPaymentSub := func() (event.Subscription, error) {
		log.Debug().
			Str("payments->connect", "Payment").
			Msgf("")
		return sdk.Contracts.Payments.WatchPayment(
			&bind.WatchOpts{Start: &blockNumber, Context: ctx},
			p.paymentChan,
		)
	}

	paymentSub, err = connectPaymentSub()
	if err != nil {
		log.Error().Err(err).Msgf("subscribe to payment requests failed")
		return err
	}

	go func() {
		<-ctx.Done()
		paymentSub.Unsubscribe()
	}()

	for {
		select {
		case event := <-p.paymentChan:
			log.Debug().
				Str("payments->event", "Payment").
				Msgf("%+v", event)
			for _, handler := range p.paymentSubs {
				go handler(*event)
			}
		case err := <-paymentSub.Err():
			paymentSub.Unsubscribe()
			paymentSub, err = connectPaymentSub()
			if err != nil {
				return err
			}
		}
	}
}

func (p *PaymentEventChannels) SubscribePayment(handler func(payments.PaymentsPayment)) {
	p.paymentSubs = append(p.paymentSubs, handler)
}
