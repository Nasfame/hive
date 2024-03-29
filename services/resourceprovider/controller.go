package resourceprovider

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/CoopHive/hive/pkg/dto"
	"github.com/CoopHive/hive/pkg/executor"
	"github.com/CoopHive/hive/pkg/http"
	"github.com/CoopHive/hive/pkg/module"
	"github.com/CoopHive/hive/pkg/system"
	"github.com/CoopHive/hive/pkg/web3"
	"github.com/CoopHive/hive/pkg/web3/bindings/storage"
	"github.com/CoopHive/hive/services/dealmaker"
	solver2 "github.com/CoopHive/hive/services/solver/solver"
	"github.com/CoopHive/hive/services/solver/solver/store"
)

type ResourceProviderController struct {
	solverClient *solver2.SolverClient
	options      ResourceProviderOptions
	web3SDK      *web3.Web3SDK
	web3Events   *web3.EventChannels
	loop         *system.ControlLoop
	log          *system.ServiceLogger
	executor     executor.Executor
	// keep track of which jobs are running
	// this is because no remote state will change
	// whilst we are actually running a job
	runningJobsMutex sync.RWMutex
	runningJobs      map[string]bool

	dealmakerService *dealmaker.Service
	dealContainers   map[string]*dto.DealContainer
	once             sync.Once
}

// the background "even if we have not heard of an event" loop
// i.e. things will not wait 10 seconds - the control loop
// reacts to events in the system - this 10 second background
// loop is just for in case we miss any events
const CONTROL_LOOP_INTERVAL = 10 * time.Second

func NewResourceProviderController(
	options ResourceProviderOptions,
	web3SDK *web3.Web3SDK,
	executor executor.Executor,
	dealmakerService *dealmaker.Service,
) (*ResourceProviderController, error) {
	// we know the address of the solver but what is it's url?
	solverUrl, err := web3SDK.GetSolverUrl(options.Offers.Services.Solver)
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

	controller := &ResourceProviderController{
		solverClient:     solverClient,
		options:          options,
		web3SDK:          web3SDK,
		web3Events:       web3.NewEventChannels(),
		log:              system.NewServiceLogger(system.ResourceProviderService),
		executor:         executor,
		runningJobs:      map[string]bool{},
		dealmakerService: dealmakerService,
		dealContainers:   map[string]*dto.DealContainer{},
	}
	return controller, nil
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
func (controller *ResourceProviderController) subscribeToSolver() error {
	controller.solverClient.SubscribeEvents(func(ev solver2.SolverEvent) {
		// we need to agree to the deal now we've heard about it
		if ev.EventType == solver2.DealAdded {
			if ev.Deal == nil {
				controller.log.Error("solver event", fmt.Errorf("RP received nil deal"))
				return
			}

			// check if this deal is for us
			if ev.Deal.ResourceProvider != controller.web3SDK.GetAddress().String() {
				return
			}

			solver2.ServiceLogSolverEvent(system.ResourceProviderService, ev)

			// trigger the solver
			controller.loop.Trigger()
		}
	})
	return nil
}

func (controller *ResourceProviderController) subscribeToWeb3() error {
	controller.web3Events.Storage.SubscribeDealStateChange(func(ev storage.StorageDealStateChange) {
		deal, err := controller.solverClient.GetDeal(ev.DealId)
		if err != nil {
			controller.log.Error("error getting deal", err)
			return
		}
		if deal.ResourceProvider != controller.web3SDK.GetAddress().String() {
			return
		}
		// FIXME:the log caller is incorrect
		controller.log.Trace("StorageDealStateChange", dto.GetAgreementStateString(ev.State))
		system.DumpObjectDebug(ev)
		controller.loop.Trigger()
	})
	return nil
}

func (controller *ResourceProviderController) Start(ctx context.Context, cm *system.CleanupManager) chan error {
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
		system.ResourceProviderService,
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

func (controller *ResourceProviderController) solve() error {
	controller.log.Debug("solving", "")

	// if the solver does not know about resource offers
	// that we have - we should post them to the solver
	err := controller.ensureResourceOffers()
	if err != nil {
		return err
	}

	// if there are deals that have been matched and we have not agreed
	// then we should agree to them
	err = controller.agreeToMatchedDeals()
	if err != nil {
		return err
	}

	// if there are jobs that have had both sides agree then we should run the job
	err = controller.runJobs()
	if err != nil {
		return err
	}

	return nil
}

/*
 *
 *
 *

 Ensure resource offers

 *
 *
 *
*/

/*
Ensure resource offers are posted to the solve
*/

func (controller *ResourceProviderController) getResourceOffer(index int, spec dto.MachineSpec) dto.ResourceOffer {
	return dto.ResourceOffer{
		// assign CreatedAt to the current millisecond timestamp
		CreatedAt:        int(time.Now().UnixNano() / int64(time.Millisecond)),
		ResourceProvider: controller.web3SDK.GetAddress().String(),
		Index:            index,
		Spec:             spec,
		Modules:          controller.options.Offers.Modules,
		Mode:             controller.options.Offers.Mode,
		DefaultPricing:   controller.options.Offers.DefaultPricing,
		DefaultTimeouts:  controller.options.Offers.DefaultTimeouts,
		ModulePricing:    map[string]dto.DealPricing{},
		ModuleTimeouts:   map[string]dto.DealTimeouts{},
		Services:         controller.options.Offers.Services,
	}
}

func (controller *ResourceProviderController) ensureResourceOffers() error {
	// load the resource offers that are currently active and so should not be replaced
	activeResourceOffers, err := controller.solverClient.GetResourceOffers(store.GetResourceOffersQuery{
		ResourceProvider: controller.web3SDK.GetAddress().String(),
		Active:           true,
	})
	if err != nil {
		return err
	}

	// create a map of the ids of resource offers we have
	// this will allow us to check if we need to create a new one
	// or update an existing one - we use the "index" because
	// the id's are changing because of the timestamps
	existingResourceOffersMap := map[int]dto.ResourceOfferContainer{}
	for _, existingResourceOffer := range activeResourceOffers {
		existingResourceOffersMap[existingResourceOffer.ResourceOffer.Index] = existingResourceOffer
	}

	addResourceOffers := []dto.ResourceOffer{}

	// map over the specs we have in the config
	for index, spec := range controller.options.Offers.Specs {

		// check if the resource offer already exists
		// if it does then we need to update it
		// if it doesn't then we need to add it
		_, ok := existingResourceOffersMap[index]
		if !ok {
			addResourceOffers = append(addResourceOffers, controller.getResourceOffer(index, spec))
		}
	}

	// add the resource offers we need to add
	for _, resourceOffer := range addResourceOffers {
		controller.log.Info("add resource offer", resourceOffer)
		_, err := controller.solverClient.AddResourceOffer(resourceOffer)
		if err != nil {
			controller.log.Error("error adding resource offer", err)
			return err
		}
	}

	return err
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
func (controller *ResourceProviderController) agreeToMatchedDeals() error {
	// load all deals that are in DealAgreed state and are for us
	matchedDeals, err := controller.solverClient.GetDealsWithFilter(
		store.GetDealsQuery{
			ResourceProvider: controller.web3SDK.GetAddress().String(),
			State:            "DealNegotiating",
		},
		// if we have already submitted an agree tx then don't do it again
		func(dealContainer dto.DealContainer) bool {
			return dealContainer.Transactions.ResourceProvider.Agree == ""
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
		// controller.log.Debug("dealContainer", dealContainer)
		dealContainers[dealContainer.ID] = &dealContainer
		go controller.dealmakerService.DealMatched(dealContainer.ID)
	}

	go controller.once.Do(func() {
		controller.dealmakerService.DealsAgreed(func(dealID string) error {
			controller.log.Debug("deal agreed ", dealID)
			return controller.agreeDeal(controller.dealContainers[dealID])
		})
	})

	return err

}

func (controller *ResourceProviderController) agreeDeal(dealContainer *dto.DealContainer) error {
	controller.log.Info("agree", dealContainer)

	txHash, err := controller.web3SDK.Agree(dealContainer.Deal)
	if err != nil {
		// TODO: we need a way of deciding based on certain classes of error what happens
		// some will be retryable - otherwise will be fatal
		// we need a way to exit a job loop as a baseline
		controller.log.Error("error calling agree tx for deal", err)
		return err
	}
	controller.log.Info("agree tx", txHash)

	// we have agreed to the deal so we need to update the tx in the solver
	_, err = controller.solverClient.UpdateTransactionsResourceProvider(dealContainer.ID, dto.DealTransactionsResourceProvider{
		Agree: txHash,
	})
	if err != nil {
		// TODO: we need a way of deciding based on certain classes of error what happens
		// some will be retryable - otherwise will be fatal
		// we need a way to exit a job loop as a baseline
		controller.log.Error("error adding agree tx hash for deal", err)
		return err
	}
	controller.log.Info("updated deal with agree tx", txHash)
	return nil
}

/*
 *
 *
 *

 Run jobs

 *
 *
 *
*/

func (controller *ResourceProviderController) runJobs() error {
	agreedDeals, err := controller.solverClient.GetDealsWithFilter(
		store.GetDealsQuery{
			ResourceProvider: controller.web3SDK.GetAddress().String(),
			State:            "DealAgreed",
		},
		// this is where the solver has found us a match and we need to agree to it
		func(dealContainer dto.DealContainer) bool {
			controller.runningJobsMutex.RLock()
			defer controller.runningJobsMutex.RUnlock()
			_, ok := controller.runningJobs[dealContainer.ID]
			return !ok
		},
	)
	if err != nil {
		return err
	}

	if len(agreedDeals) <= 0 {
		return nil
	}

	// TODO - we are relying on the rate at which we post resource offers
	// as our capacity manager right now
	// this will work because we only post resource offers as fast we handle jobs
	// but it would be worth putting some kind of queue here that is also aware
	// of the underlying capacity of the machine

	// map over the deals and run them
	for _, dealContainer := range agreedDeals {
		func() {
			controller.runningJobsMutex.Lock()
			defer controller.runningJobsMutex.Unlock()

			controller.runningJobs[dealContainer.ID] = true
		}()

		go controller.runJob(dealContainer)
	}

	return err
}

// this is run in it's own go-routine
// we've already updated controller.runningJobs so we know this will only
// run once
func (controller *ResourceProviderController) runJob(deal dto.DealContainer) {
	controller.log.Info("run job", deal)
	result := dto.Result{
		DealID: deal.ID,
		Error:  "",
	}
	err := func() error {
		controller.log.Info("loading module", "")
		module, err := module.LoadModule(deal.Deal.JobOffer.Module, deal.Deal.JobOffer.Inputs)
		if err != nil {
			controller.log.Error("error loading module", err)
			return fmt.Errorf("error loading module: %s", err.Error())
		}
		controller.log.Info("module loaded", module)
		executorResult, err := controller.executor.RunJob(deal, *module)
		if err != nil {
			controller.log.Error("error running job", err)
			return fmt.Errorf("error running job: %s", err.Error())
		}
		result.InstructionCount = uint64(executorResult.InstructionCount)
		result.DataID = executorResult.ResultsCID
		controller.log.Info("got result", result)

		controller.log.Info(fmt.Sprintf("uploading results: %s %s %s", deal.ID, executorResult.ResultsDir, executorResult.ResultsCID), executorResult.ResultsDir)

		_, err = controller.solverClient.UploadResultFiles(deal.ID, executorResult.ResultsDir)
		if err != nil {
			controller.log.Error("error uploading results", err)
			return fmt.Errorf("error uploading results: %s", err.Error())
		}

		return nil
	}()

	// if this error is defined then it is probably the fault of the job not us
	// and we expect a mediator to get the same error
	if err != nil {
		result.Error = err.Error()
	}

	// the tarball of the results has been uploaded
	// now let's post the result data itself to the solver
	// then we will post the results on-chain
	createdResult, err := controller.solverClient.AddResult(result)
	if err != nil {
		// TODO: what should we do here?
		// the current path would be the post results times out
		// and the JC can claim a refund
		// but it's not really the fault of the RP that the solver refused to upload the results
		controller.log.Error("error posting result", err)
		return
	}

	txHash, err := controller.web3SDK.AddResult(
		deal.Deal.ID,
		createdResult.ID,
		createdResult.DataID,
		result.InstructionCount,
	)
	if err != nil {
		controller.log.Error("error calling add result tx for job", err)
		return
	}

	_, err = controller.solverClient.UpdateTransactionsResourceProvider(deal.ID, dto.DealTransactionsResourceProvider{
		AddResult: txHash,
	})
	if err != nil {
		// TODO: we need a way of deciding based on certain classes of error what happens
		// some will be retryable - otherwise will be fatal
		// we need a way to exit a job loop as a baseline
		controller.log.Error("error adding add result tx hash for deal", err)
		return
	}
}
