package jobCreatorService

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/CoopHive/hive/pkg/dto"
	"github.com/CoopHive/hive/pkg/http"
	"github.com/CoopHive/hive/pkg/system"
	"github.com/CoopHive/hive/pkg/web3"
	"github.com/CoopHive/hive/pkg/web3/bindings/storage"
	"github.com/CoopHive/hive/services/dealmaker"
	solver2 "github.com/CoopHive/hive/services/solver/solver"
	"github.com/CoopHive/hive/services/solver/solver/store"
)

type JobOfferSubscriber func(offer dto.JobOfferContainer)

type JobCreatorController struct {
	solverClient          *solver2.SolverClient
	options               JobCreatorOptions
	web3SDK               *web3.Web3SDK
	web3Events            *web3.EventChannels
	loop                  *system.ControlLoop
	log                   *system.ServiceLogger
	jobOfferSubscriptions []JobOfferSubscriber

	dealmakerService *dealmaker.Service

	dealContainers map[string]*dto.DealContainer

	once sync.Once
}

// the background "even if we have not heard of an event" loop
// i.e. things will not wait 10 seconds - the control loop
// reacts to events in the system - this 10 second background
// loop is just for in case we miss any events
const CONTROL_LOOP_INTERVAL = 10 * time.Second

func NewJobCreatorController(
	options JobCreatorOptions,
	web3SDK *web3.Web3SDK,
	dealmakerService *dealmaker.Service,
) (*JobCreatorController, error) {
	// we know the address of the solver but what is it's url?
	solverUrl, err := web3SDK.GetSolverUrl(options.Offer.Services.Solver)
	if err != nil {
		return nil, err
	}

	solverClient, err := solver2.NewSolverClient(http.ClientOptions{
		URL:        solverUrl,
		PrivateKey: options.Web3.PrivateKey,
	})
	if err != nil {
		return nil, err
	}

	controller := &JobCreatorController{
		solverClient:          solverClient,
		options:               options,
		web3SDK:               web3SDK,
		web3Events:            web3.NewEventChannels(),
		log:                   system.NewServiceLogger(system.JobCreatorService),
		jobOfferSubscriptions: []JobOfferSubscriber{},
		dealmakerService:      dealmakerService,
		dealContainers:        map[string]*dto.DealContainer{},
	}
	return controller, nil
}

/*
 *
 *
 *

 Public API

 *
 *
 *
*/

func (controller *JobCreatorController) AddJobOffer(offer dto.JobOffer) (dto.JobOfferContainer, error) {
	controller.log.Debug("add job offer", offer)
	return controller.solverClient.AddJobOffer(offer)
}

func (controller *JobCreatorController) SubscribeToJobOfferUpdates(sub JobOfferSubscriber) {
	controller.jobOfferSubscriptions = append(controller.jobOfferSubscriptions, sub)
}

/*
*
*
*

	Setup

*
*
*
*/
func (controller *JobCreatorController) subscribeToSolver() error {
	controller.solverClient.SubscribeEvents(func(ev solver2.SolverEvent) {
		switch ev.EventType {
		case solver2.DealAdded:
			if ev.Deal == nil {
				controller.log.Error("solver event", fmt.Errorf("JC received nil deal"))
				return
			}

			// check if this deal is for us
			if ev.Deal.JobCreator != controller.web3SDK.GetAddress().String() {
				return
			}

			solver2.ServiceLogSolverEvent(system.JobCreatorService, ev)

			// trigger the solver
			controller.loop.Trigger()
		case solver2.JobOfferStateUpdated:
			if ev.JobOffer == nil {
				controller.log.Error("solver event", fmt.Errorf("RP received nil job offer"))
				return
			}
			for _, sub := range controller.jobOfferSubscriptions {
				go sub(*ev.JobOffer)
			}
		}
	})
	return nil
}

func (controller *JobCreatorController) subscribeToWeb3() error {
	controller.web3Events.Storage.SubscribeDealStateChange(func(ev storage.StorageDealStateChange) {
		deal, err := controller.solverClient.GetDeal(ev.DealId)
		if err != nil {
			controller.log.Error("error getting deal", err)
			return
		}
		if deal.JobCreator != controller.web3SDK.GetAddress().String() {
			return
		}
		controller.log.Debug("StorageDealStateChange", dto.GetAgreementStateString(ev.State))
		system.DumpObjectDebug(ev)
		controller.loop.Trigger()
	})
	return nil
}

func (controller *JobCreatorController) Start(ctx context.Context, cm *system.CleanupManager) chan error {
	errorChan := make(chan error)
	err := controller.subscribeToSolver()
	if err != nil {
		errorChan <- err
		return errorChan
	}
	err = controller.subscribeToWeb3()
	if err != nil {
		errorChan <- err
		return errorChan
	}

	// this connects the websocket client
	err = controller.solverClient.Start(ctx, cm)
	if err != nil {
		errorChan <- err
		return errorChan
	}
	err = controller.web3Events.Start(controller.web3SDK, ctx, cm)
	if err != nil {
		errorChan <- err
		return errorChan
	}

	controller.loop = system.NewControlLoop(
		system.JobCreatorService,
		ctx,
		CONTROL_LOOP_INTERVAL,
		func() error {
			err := controller.solve()
			if err != nil {
				errorChan <- err
			}
			return err
		},
	)

	err = controller.loop.Start(true)
	if err != nil {
		errorChan <- err
		return errorChan
	}

	return errorChan
}

/*
 *
 *
 *

 Solve

 *
 *
 *
*/

func (controller *JobCreatorController) solve() error {
	controller.log.Debug("solving", "")
	err := controller.agreeToMatchedDeals()
	if err != nil {
		return err
	}
	err = controller.checkResults()
	if err != nil {
		return err
	}
	return nil
}

/*
 *
 *
 *

 Agree to deals

 *
 *
 *
*/

// list the deals we have been assigned to that we have not yet posted and agree tx to the contract for
func (controller *JobCreatorController) agreeToMatchedDeals() error {

	// load the deals that are in DealNegotiating
	// and do not have a TransactionsResourceProvider.Agree tx
	matchedDeals, err := controller.solverClient.GetDealsWithFilter(
		store.GetDealsQuery{
			JobCreator: controller.web3SDK.GetAddress().String(),
			State:      "DealNegotiating",
		},
		// this is where the solver has found us a match and we need to agree to it
		func(dealContainer dto.DealContainer) bool {
			return dealContainer.Transactions.JobCreator.Agree == ""
		},
	)
	if err != nil {
		return err
	}
	if len(matchedDeals) == 0 {
		controller.log.Debug("matchedDeals", "found no deals")
		return nil
	}

	dealContainers := controller.dealContainers

	// map over the deals and agree to them
	for _, dealContainer := range matchedDeals {
		controller.log.Debug("dealContainer", dealContainer)
		dealContainers[dealContainer.ID] = &dealContainer
		go controller.dealmakerService.DealMatched(dealContainer.ID)
	}
	go controller.once.Do(func() {
		controller.dealmakerService.DealsAgreed(func(dealID string) {
			controller.log.Debug("deal agreed ", dealID)
			controller.agreeDeal(controller.dealContainers[dealID])
		})
	})

	return err
}

func (controller *JobCreatorController) agreeDeal(dealContainer *dto.DealContainer) {
	controller.log.Debug("[controller] agree", dealContainer)
	txHash, err := controller.web3SDK.Agree(dealContainer.Deal)
	if err != nil {
		// TODO: error handling - is it terminal or retryable?
		controller.log.Error("error calling agree tx for deal", err)
		return
	}
	controller.log.Debug("agree tx", txHash)

	// we have agreed to the deal so we need to update the tx in the solver
	_, err = controller.solverClient.UpdateTransactionsJobCreator(dealContainer.ID, dto.DealTransactionsJobCreator{
		Agree: txHash,
	})
	if err != nil {
		// TODO: error handling - is it terminal or retryable?
		controller.log.Error("error adding agree tx hash for deal", err)
		return
	}
	controller.log.Debug("updated deal with agree tx", txHash)
}

// list the deals that have results posted but we have not yet checked
// we do this synchronously to prevent us racing with large result sets
// also we are the client so have a lower chance of there being a chunky backlog
func (controller *JobCreatorController) checkResults() error {
	// load all deals in ResultsSubmitted state and don't have either results checked or accepted txs
	completedDeals, err := controller.solverClient.GetDealsWithFilter(
		store.GetDealsQuery{
			JobCreator: controller.web3SDK.GetAddress().String(),
			State:      "ResultsSubmitted",
		},
		// this is where the solver has found us a match and we need to agree to it
		func(dealContainer dto.DealContainer) bool {
			return dealContainer.Transactions.JobCreator.AcceptResult == "" && dealContainer.Transactions.JobCreator.CheckResult == ""
		},
	)
	if err != nil {
		return err
	}
	if len(completedDeals) <= 0 {
		return nil
	}

	for _, dealContainer := range completedDeals {
		result, err := controller.solverClient.GetResult(dealContainer.ID)
		controller.log.Debug("result", result)

		if err != nil || result.Error != "" {
			// there is an error with the job
			// accept anyway
			// TODO: trigger mediation here
			// controller.checkResult(dealContainer) //TODO:
			err = controller.acceptResult(dealContainer)
			controller.log.Error("result errored with %v", err)
		} else {
			err = controller.downloadResult(dealContainer)
		}
	}

	return err
}

func (controller *JobCreatorController) downloadResult(dealContainer dto.DealContainer) error {
	err := controller.solverClient.DownloadResultFiles(dealContainer.ID, solver2.GetDownloadsFilePath(dealContainer.ID))
	if err != nil {
		// TODO: jc didn't accept or deny results
		return fmt.Errorf("error downloading results for deal: %w", err)
	}

	controller.log.Debug("Downloaded results for job", solver2.GetDownloadsFilePath(dealContainer.ID))

	// TODO: activate the mediation check here
	controller.acceptResult(dealContainer)

	// work out if we should check or accept the results
	// if controller.options.Mediation.CheckResultsPercentage >= rand.Intn(100) {
	// 	err = controller.checkResult(dealContainer)

	// 	if err != nil {
	// 		// TODO: error handling - is it terminal or retryable?
	// 		controller.log.Error("error checking deal results", err)
	// 		return nil
	// 	}

	// 	controller.log.Debug("Checked results for job", dealContainer.ID)
	// } else {
	// 	err = controller.acceptResult(dealContainer)

	// 	if err != nil {
	// 		// TODO: error handling - is it terminal or retryable?
	// 		controller.log.Error("error accepting deal results", err)
	// 		return nil
	// 	}

	// 	controller.log.Debug("Accepted results for job", dealContainer.ID)
	// }
	return nil
}

func (controller *JobCreatorController) acceptResult(deal dto.DealContainer) error {
	controller.log.Debug("Accepting results for job", deal.ID)
	txHash, err := controller.web3SDK.AcceptResult(deal.ID)
	if err != nil {
		return fmt.Errorf("error calling accept result tx for deal: %w", err)
	}
	controller.log.Debug("accept result tx", txHash)

	// we have agreed to the deal so we need to update the tx in the solver
	_, err = controller.solverClient.UpdateTransactionsJobCreator(deal.ID, dto.DealTransactionsJobCreator{
		AcceptResult: txHash,
	})
	if err != nil {
		return fmt.Errorf("error adding AcceptResult tx hash for deal: %s", err.Error())
	}
	return nil
}

func (controller *JobCreatorController) checkResult(deal dto.DealContainer) error {
	controller.log.Debug("Checking results for job", deal.ID)
	txHash, err := controller.web3SDK.CheckResult(deal.ID)
	if err != nil {
		return fmt.Errorf("error calling check result tx for deal: %s", err.Error())
	}
	controller.log.Debug("check result tx", txHash)

	// we have agreed to the deal so we need to update the tx in the solver
	_, err = controller.solverClient.UpdateTransactionsJobCreator(deal.ID, dto.DealTransactionsJobCreator{
		CheckResult: txHash,
	})
	if err != nil {
		return fmt.Errorf("error adding CheckResult tx hash for deal: %s", err.Error())
	}
	return nil
}
