package web3

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/event"
	"github.com/rs/zerolog/log"

	"github.com/CoopHive/hive/pkg/system"
	"github.com/CoopHive/hive/pkg/web3/bindings/mediation"
)

type MediationEventChannels struct {
	mediationRequestedChan chan *mediation.MediationMediationRequested
	mediationRequestedSubs []func(mediation.MediationMediationRequested)
}

func NewMediationEventChannels() *MediationEventChannels {
	return &MediationEventChannels{
		mediationRequestedChan: make(chan *mediation.MediationMediationRequested),
		mediationRequestedSubs: []func(mediation.MediationMediationRequested){},
	}
}

func (m *MediationEventChannels) Start(
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

	var mediationRequestedSub event.Subscription

	connectMediationRequestedSub := func() (event.Subscription, error) {
		log.Debug().
			Str("mediation->connect", "MediationRequested").
			Msgf("")
		return sdk.Contracts.Mediation.WatchMediationRequested(
			&bind.WatchOpts{Start: &blockNumber, Context: ctx},
			m.mediationRequestedChan,
		)
	}

	mediationRequestedSub, err = connectMediationRequestedSub()
	if err != nil {
		log.Error().Err(err).Msgf("subscribe to mediation requests failed")
		return err
	}

	go func() {
		<-ctx.Done()
		mediationRequestedSub.Unsubscribe()
	}()

	for {
		select {
		case event := <-m.mediationRequestedChan:
			log.Debug().
				Str("mediation->event", "MediationRequested").
				Msgf("%+v", event)
			for _, handler := range m.mediationRequestedSubs {
				go handler(*event)
			}
		case err := <-mediationRequestedSub.Err():
			mediationRequestedSub.Unsubscribe()
			mediationRequestedSub, err = connectMediationRequestedSub()
			if err != nil {
				return err
			}
		}
	}
}

func (m *MediationEventChannels) SubscribeMediationRequested(handler func(mediation.MediationMediationRequested)) {
	m.mediationRequestedSubs = append(m.mediationRequestedSubs, handler)
}
