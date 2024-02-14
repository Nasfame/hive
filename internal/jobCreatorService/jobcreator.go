package jobCreatorService

import (
	"context"

	"github.com/CoopHive/hive/pkg/dto"
	"github.com/CoopHive/hive/pkg/system"
	"github.com/CoopHive/hive/pkg/web3"
	"github.com/CoopHive/hive/services/dealmaker"
)

type JobCreatorMediationOptions struct {
	// out of 100 chance we will check results
	CheckResultsPercentage int
}

type JobCreatorOfferOptions struct {
	// the module that is wanting to be run
	// this contains the spec that is required to run the module
	Module dto.ModuleConfig
	// the required spec hoisted from the module
	Spec dto.MachineSpec
	// this will normally be MarketPrice for JC's
	Mode dto.PricingMode
	// this is so clients can put limit orders for jobs
	// and the solver will match as soon as a resource offer
	// is added that matches the bid
	Pricing dto.DealPricing
	// the timeouts we are offering with the deal
	Timeouts dto.DealTimeouts
	// the inputs to the module
	Inputs map[string]string
	// which mediators and directories this RP will trust
	Services dto.ServiceConfig
}

type JobCreatorOptions struct {
	Mediation JobCreatorMediationOptions
	Offer     JobCreatorOfferOptions
	Web3      web3.Web3Options
	Dealer    string
}

type JobCreator struct {
	web3SDK    *web3.Web3SDK
	options    JobCreatorOptions
	controller *JobCreatorController
}

func NewJobCreator(
	options JobCreatorOptions,
	web3SDK *web3.Web3SDK,
	dealmakerService *dealmaker.Service,
) (*JobCreator, error) {
	controller, err := NewJobCreatorController(options, web3SDK, dealmakerService)
	if err != nil {
		return nil, err
	}
	jc := &JobCreator{
		web3SDK,
		options,
		controller,
	}
	return jc, nil
}

func (jobCreator *JobCreator) Start(ctx context.Context, cm *system.CleanupManager) chan error {
	return jobCreator.controller.Start(ctx, cm)
}

func (jobCreator *JobCreator) GetJobOfferFromOptions(options JobCreatorOfferOptions) (dto.JobOffer, error) {
	return getJobOfferFromOptions(options, jobCreator.web3SDK.GetAddress().String())
}

// adds the job offer to the solver
func (jobCreator *JobCreator) AddJobOffer(offer dto.JobOffer) (dto.JobOfferContainer, error) {
	return jobCreator.controller.AddJobOffer(offer)
}

func (jobCreator *JobCreator) SubscribeToJobOfferUpdates(sub JobOfferSubscriber) {
	jobCreator.controller.SubscribeToJobOfferUpdates(sub)
}

func (jobCreator *JobCreator) GetResult(dealId string) (dto.Result, error) {
	return jobCreator.controller.solverClient.GetResult(dealId)
}
