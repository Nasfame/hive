package jobCreatorService

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/rs/zerolog/log"

	"github.com/CoopHive/hive/config"
	"github.com/CoopHive/hive/enums"
	"github.com/CoopHive/hive/pkg/dto"
	"github.com/CoopHive/hive/pkg/system"
	"github.com/CoopHive/hive/pkg/web3"
	jobcreatorweb3 "github.com/CoopHive/hive/pkg/web3/bindings/jobcreator"
	"github.com/CoopHive/hive/services/dealmaker"
)

type OnChainJobCreator struct {
	web3SDK       *web3.Web3SDK
	web3Events    *web3.EventChannels
	options       JobCreatorOptions
	controller    *JobCreatorController
	onChainJobIDs map[string]uint64
}

func NewOnChainJobCreator(
	options JobCreatorOptions,
	web3SDK *web3.Web3SDK,
	dealmakerService *dealmaker.Service,
) (*OnChainJobCreator, error) {
	controller, err := NewJobCreatorController(options, web3SDK, dealmakerService)
	if err != nil {
		return nil, err
	}
	jc := &OnChainJobCreator{
		controller:    controller,
		options:       options,
		web3SDK:       web3SDK,
		web3Events:    web3.NewEventChannels(),
		onChainJobIDs: map[string]uint64{},
	}
	return jc, nil
}

func (jobCreator *OnChainJobCreator) Start(ctx context.Context, cm *system.CleanupManager) chan error {
	errorChan := jobCreator.controller.Start(ctx, cm)

	jobPrice := config.Conf.GetFloat64(enums.JOB_PRICE)

	// TODO: work out how to do dynamic pricing
	tx, err := jobCreator.web3SDK.Contracts.JobCreator.SetRequiredDeposit(jobCreator.web3SDK.TransactOpts, web3.EtherToWei(jobPrice))

	_, err = jobCreator.web3SDK.WaitTx(ctx, tx, err)
	if err != nil {
		log.Debug().Err(err).Msgf("jobCreator Start")
		errorChan <- err
		return errorChan
	}

	err = jobCreator.web3Events.Start(jobCreator.web3SDK, ctx, cm)
	if err != nil {
		errorChan <- err
		return errorChan
	}

	jobCreator.controller.SubscribeToJobOfferUpdates(func(evOffer dto.JobOfferContainer) {
		if evOffer.State != dto.GetAgreementStateIndex("ResultsAccepted") {
			return
		}

		onChainID, ok := jobCreator.onChainJobIDs[evOffer.DealID]
		if !ok {
			return
		}

		result, err := jobCreator.GetResult(evOffer.DealID)
		if err != nil {
			return
		}

		fmt.Printf("result --------------------------------------\n")
		spew.Dump(result)
		spew.Dump(int64(onChainID))

		tx, err := jobCreator.web3SDK.Contracts.JobCreator.SubmitResults(jobCreator.web3SDK.TransactOpts, big.NewInt(int64(onChainID)), evOffer.DealID, result.DataID)

		_, err = jobCreator.web3SDK.WaitTx(ctx, tx, err)
		if err != nil {
			log.Debug().Err(err).Msgf("jc Start")
			return
		}
	})

	jobCreator.web3Events.JobCreator.SubscribeJobAdded(func(ev jobcreatorweb3.JobcreatorJobAdded) {

		// first we need to move the tokens into our account
		tx, err := jobCreator.web3SDK.Contracts.Token.TransferFrom(jobCreator.web3SDK.TransactOpts, ev.Payee, jobCreator.web3SDK.GetAddress(), web3.EtherToWei(jobPrice))

		_, err = jobCreator.web3SDK.WaitTx(ctx, tx, err)
		if err != nil {
			log.Debug().Err(err).Msg("error creating job offer")
			return
		}

		options := jobCreator.options.Offer
		options.Module.Name = ev.Module
		inputs := map[string]string{}
		for _, input := range ev.Inputs {
			parts := strings.Split(input, "=")
			if len(parts) == 2 {
				inputs[parts[0]] = parts[1]
			}
		}
		options.Inputs = inputs
		offer, err := getJobOfferFromOptions(options, jobCreator.web3SDK.GetAddress().String())
		if err != nil {
			fmt.Printf("error creating job offer: %s\n", err.Error())
			return
		}

		container, err := jobCreator.controller.AddJobOffer(offer)
		if err != nil {
			fmt.Printf("error creating job offer: %s\n", err.Error())
			return
		}

		jobCreator.onChainJobIDs[container.DealID] = ev.Id.Uint64()
	})

	return errorChan
}

func (jobCreator *OnChainJobCreator) GetJobOfferFromOptions(options JobCreatorOfferOptions) (dto.JobOffer, error) {
	return getJobOfferFromOptions(options, jobCreator.web3SDK.GetAddress().String())
}

// adds the job offer to the solver
func (jobCreator *OnChainJobCreator) AddJobOffer(offer dto.JobOffer) (dto.JobOfferContainer, error) {
	return jobCreator.controller.AddJobOffer(offer)
}

func (jobCreator *OnChainJobCreator) SubscribeToJobOfferUpdates(sub JobOfferSubscriber) {
	jobCreator.controller.SubscribeToJobOfferUpdates(sub)
}

func (jobCreator *OnChainJobCreator) GetResult(dealId string) (dto.Result, error) {
	return jobCreator.controller.solverClient.GetResult(dealId)
}
