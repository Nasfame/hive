package web3

import (
	"context"

	"github.com/CoopHive/hive/pkg/system"
	"github.com/CoopHive/hive/utils"
)

type EventChannels struct {
	Token       *TokenEventChannels
	Payment     *PaymentEventChannels
	Storage     *StorageEventChannels
	JobCreator  *JobCreatorEventChannels
	Mediation   *MediationEventChannels
	collections []EventChannelCollection
}

func NewEventChannels() *EventChannels {
	tokenChannels := NewTokenEventChannels()
	paymentChannels := NewPaymentEventChannels()
	storageChannels := NewStorageEventChannels()
	jobCreatorChannels := NewJobCreatorEventChannels()
	mediationChannels := NewMediationEventChannels()
	collections := []EventChannelCollection{
		tokenChannels,
		paymentChannels,
		storageChannels,
		jobCreatorChannels,
		mediationChannels,
	}
	return &EventChannels{
		Token:       tokenChannels,
		Payment:     paymentChannels,
		Storage:     storageChannels,
		JobCreator:  jobCreatorChannels,
		Mediation:   mediationChannels,
		collections: collections,
	}
}

func (eventChannels *EventChannels) Start(
	sdk *Web3SDK,
	ctx context.Context,
	cm *system.CleanupManager,
) error {
	utils.PanicOnHTTPUrl(sdk.Options.RpcURL)

	for _, collection := range eventChannels.collections {
		c := collection
		go func() {
			c.Start(sdk, ctx, cm) // TODO:
			/*	if err != nil {
				log.Error().Msgf("error starting listeners: %s", err.Error())
				panic("panic starting listeners")
			}*/
		}()
	}
	return nil
}

func eventErrorHandler(err error) {
	if err != nil {
		panic(err)
	}
}
