package resourceprovider

import (
	"context"

	"github.com/CoopHive/hive/pkg/dto"
	"github.com/CoopHive/hive/pkg/executor"
	"github.com/CoopHive/hive/pkg/executor/bacalhau"
	"github.com/CoopHive/hive/pkg/system"
	"github.com/CoopHive/hive/pkg/web3"
)

// this configures the resource offers we will keep track of
type ResourceProviderOfferOptions struct {
	// if we are configuring a single machine then
	// these values are populated by the flags
	OfferSpec dto.MachineSpec
	// we can dupliate the single spec to create a list of specs
	OfferCount int
	// this represents how many machines we will keep
	// offering to the network
	// we can configure this with a config file
	// to start with we will just add --cpu --gpu and --ram flags
	// to the resource provider CLI which constrains them to a single machine
	Specs []dto.MachineSpec
	// the list of modules we are willing to run
	// an empty list means anything
	Modules []string

	// this will normally be FixedPrice for RP's
	Mode dto.PricingMode

	// the default pricing for this resource provider
	// for all modules that don't have a specific price
	DefaultPricing  dto.DealPricing
	DefaultTimeouts dto.DealTimeouts

	// allow different pricing for different modules
	ModulePricing  map[string]dto.DealPricing
	ModuleTimeouts map[string]dto.DealTimeouts

	// which mediators and directories this RP will trust
	Services dto.ServiceConfig
}

type ResourceProviderOptions struct {
	Bacalhau bacalhau.BacalhauExecutorOptions
	Offers   ResourceProviderOfferOptions
	Web3     web3.Web3Options
}

type ResourceProvider struct {
	web3SDK    *web3.Web3SDK
	options    ResourceProviderOptions
	controller *ResourceProviderController
}

func NewResourceProvider(
	options ResourceProviderOptions,
	web3SDK *web3.Web3SDK,
	executor executor.Executor,
) (*ResourceProvider, error) {
	controller, err := NewResourceProviderController(options, web3SDK, executor)
	if err != nil {
		return nil, err
	}
	solver := &ResourceProvider{
		controller: controller,
		options:    options,
		web3SDK:    web3SDK,
	}
	return solver, nil
}

func (resourceProvider *ResourceProvider) Start(ctx context.Context, cm *system.CleanupManager) chan error {
	return resourceProvider.controller.Start(ctx, cm)
}
