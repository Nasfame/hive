package web3

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/event"
	"github.com/rs/zerolog/log"

	"github.com/CoopHive/hive/pkg/system"
	"github.com/CoopHive/hive/pkg/web3/bindings/jobcreator"
)

type JobCreatorEventChannels struct {
	jobAddedChan chan *jobcreator.JobcreatorJobAdded
	jobAddedSubs []func(jobcreator.JobcreatorJobAdded)
}

func NewJobCreatorEventChannels() *JobCreatorEventChannels {
	return &JobCreatorEventChannels{
		jobAddedChan: make(chan *jobcreator.JobcreatorJobAdded),
		jobAddedSubs: []func(jobcreator.JobcreatorJobAdded){},
	}
}

func (s *JobCreatorEventChannels) Start(
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

	var jobAddedSub event.Subscription

	connectJobAddedSub := func() (event.Subscription, error) {
		log.Debug().
			Str("jobcreator->connect", "JobAdded").
			Msgf("")
		return sdk.Contracts.JobCreator.WatchJobAdded(
			&bind.WatchOpts{Start: &blockNumber, Context: ctx},
			s.jobAddedChan,
		)
	}

	jobAddedSub, err = connectJobAddedSub()
	if err != nil {
		log.Error().Err(err).Msgf("subscribe to job request failed")
		return err
	}
	go func() {
		<-ctx.Done()
		jobAddedSub.Unsubscribe()
	}()

	for {
		select {
		case event := <-s.jobAddedChan:
			log.Debug().
				Str("storage->event", "DealStateChange").
				Msgf("%+v", event)
			for _, handler := range s.jobAddedSubs {
				go handler(*event)
			}
		case err := <-jobAddedSub.Err():
			jobAddedSub.Unsubscribe()
			jobAddedSub, err = connectJobAddedSub()
			if err != nil {
				return err
			}
		}
	}
}

func (t *JobCreatorEventChannels) SubscribeJobAdded(handler func(jobcreator.JobcreatorJobAdded)) {
	t.jobAddedSubs = append(t.jobAddedSubs, handler)
}
