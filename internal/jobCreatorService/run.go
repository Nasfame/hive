package jobCreatorService

import (
	"fmt"
	"time"

	"github.com/CoopHive/hive/pkg/dto"
	"github.com/CoopHive/hive/pkg/system"
	"github.com/CoopHive/hive/pkg/web3"
)

type RunJobResults struct {
	JobOffer dto.JobOfferContainer
	Result   dto.Result
}

func RunJob(
	ctx *system.CommandContext,
	options JobCreatorOptions,
	eventSub JobOfferSubscriber,
) (*RunJobResults, error) {
	web3SDK, err := web3.NewContractSDK(options.Web3)
	if err != nil {
		return nil, err
	}

	// create the job creator and start it's control loop
	jobCreatorService, err := NewJobCreator(options, web3SDK)
	if err != nil {
		return nil, err
	}

	jobCreatorService.SubscribeToJobOfferUpdates(eventSub)

	jobCreatorErrors := jobCreatorService.Start(ctx.Ctx, ctx.Cm)

	// let's process our options into an actual job offer
	// this will also validate the module we are asking for
	offer, err := jobCreatorService.GetJobOfferFromOptions(options.Offer)
	if err != nil {
		return nil, err
	}

	// wait a short period because we've just started the job creator service
	time.Sleep(100 * time.Millisecond)

	jobOfferContainer, err := jobCreatorService.AddJobOffer(offer)
	if err != nil {
		return nil, err
	}

	updateChan := make(chan dto.JobOfferContainer)

	jobCreatorService.SubscribeToJobOfferUpdates(func(evOffer dto.JobOfferContainer) {
		// spew.Dump(evOffer)
		if evOffer.JobOffer.ID != jobOfferContainer.ID {
			return
		}
		updateChan <- evOffer
	})

	var finalJobOffer dto.JobOfferContainer

	// now we wait on the state of the job
waitloop:
	for {
		select {
		// this means the job was cancelled
		case err := <-jobCreatorErrors:
			return nil, err
		case <-ctx.Ctx.Done():
			return nil, fmt.Errorf("job cancelled")
		case finalJobOffer = <-updateChan:
			if dto.IsTerminalAgreementState(finalJobOffer.State) {
				break waitloop
			}
		}
	}

	result, err := jobCreatorService.GetResult(finalJobOffer.DealID)
	if err != nil {
		return nil, err
	}

	return &RunJobResults{
		JobOffer: finalJobOffer,
		Result:   result,
	}, nil
}