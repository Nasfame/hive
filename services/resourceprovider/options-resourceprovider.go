package resourceprovider

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/CoopHive/hive/config"
	"github.com/CoopHive/hive/enums"
	"github.com/CoopHive/hive/pkg/dto"
	options2 "github.com/CoopHive/hive/pkg/options"
	"github.com/CoopHive/hive/pkg/system"
)

func NewResourceProviderOptions() ResourceProviderOptions {
	options := ResourceProviderOptions{
		Bacalhau: options2.GetDefaultBacalhauOptions(),
		Offers:   GetDefaultResourceProviderOfferOptions(),
		Web3:     options2.GetDefaultWeb3Options(),
		Dealer:   config.Conf.GetString(enums.DEALER),
	}
	options.Web3.Service = system.ResourceProviderService
	return options
}

func GetDefaultResourceProviderOfferOptions() ResourceProviderOfferOptions {
	return ResourceProviderOfferOptions{
		// by default let's offer 1 CPU, 0 GPU and 1GB RAM
		OfferSpec: dto.MachineSpec{
			CPU: options2.GetDefaultServeOptionInt("OFFER_CPU", 1000), //nolint:gomnd
			GPU: options2.GetDefaultServeOptionInt("OFFER_GPU", 0),    //nolint:gomnd
			RAM: options2.GetDefaultServeOptionInt("OFFER_RAM", 1024), //nolint:gomnd
		},
		OfferCount: options2.GetDefaultServeOptionInt("OFFER_COUNT", 1), //nolint:gomnd
		// this can be populated by a config file
		Specs: []dto.MachineSpec{},
		// if an RP wants to only run certain modules they list them here
		// XXX SECURITY: enforce that they are specified with specific git hashes!
		Modules: options2.GetDefaultServeOptionStringArray("OFFER_MODULES", []string{}),
		// this is the default pricing mode for an RP
		Mode: options2.GetDefaultPricingMode(dto.FixedPrice),
		// this is the default pricing for a module unless it has a specific price
		DefaultPricing:  options2.GetDefaultPricingOptions(),
		DefaultTimeouts: options2.GetDefaultTimeoutOptions(),
		// allows an RP to list specific prices for each module
		ModulePricing:  map[string]dto.DealPricing{},
		ModuleTimeouts: map[string]dto.DealTimeouts{},
		Services:       options2.GetDefaultServicesOptions(),
	}
}

func AddResourceProviderOfferCliFlags(cmd *cobra.Command, offerOptions *ResourceProviderOfferOptions) {
	cmd.PersistentFlags().IntVar(
		&offerOptions.OfferSpec.CPU, "offer-cpu", offerOptions.OfferSpec.CPU,
		`How many milli-cpus to offer the network (OFFER_CPU).`,
	)
	cmd.PersistentFlags().IntVar(
		&offerOptions.OfferSpec.GPU, "offer-gpu", offerOptions.OfferSpec.GPU,
		`How many milli-gpus to offer the network (OFFER_GPU).`,
	)
	cmd.PersistentFlags().IntVar(
		&offerOptions.OfferSpec.RAM, "offer-ram", offerOptions.OfferSpec.RAM,
		`How many megabytes of RAM to offer the network (OFFER_RAM).`,
	)
	cmd.PersistentFlags().IntVar(
		&offerOptions.OfferCount, "offer-count", offerOptions.OfferCount,
		`How many machines will we offer using the cpu, ram and gpu settings (OFFER_COUNT).`,
	)
	cmd.PersistentFlags().StringArrayVar(
		&offerOptions.Modules, "offer-modules", offerOptions.Modules,
		`The modules you are willing to run (OFFER_MODULES).`,
	)
	options2.AddPricingModeCliFlags(cmd, &offerOptions.Mode)
	options2.AddPricingCliFlags(cmd, &offerOptions.DefaultPricing)
	options2.AddTimeoutCliFlags(cmd, &offerOptions.DefaultTimeouts)
	options2.AddServicesCliFlags(cmd, &offerOptions.Services)
}

func AddResourceProviderCliFlags(cmd *cobra.Command, options *ResourceProviderOptions) {
	options2.AddBacalhauCliFlags(cmd, &options.Bacalhau)
	options2.AddWeb3CliFlags(cmd, &options.Web3)
	AddResourceProviderOfferCliFlags(cmd, &options.Offers)
	cmd.PersistentFlags().StringVar(&options.Dealer, "dealer", options.Dealer, "Choose Dealer for accepting/denying solver-matched deals")
}

func CheckResourceProviderOfferOptions(options ResourceProviderOfferOptions) error {
	// loop over all specs and add up the total number of cpus
	totalCPU := 0
	for _, spec := range options.Specs {
		totalCPU += spec.CPU
	}

	if totalCPU <= 0 {
		return fmt.Errorf("OFFER_CPU cannot be zero")
	}

	// do the same for memory
	totalRAM := 0
	for _, spec := range options.Specs {
		totalRAM += spec.RAM
	}

	if totalRAM <= 0 {
		return fmt.Errorf("OFFER_RAM cannot be zero")
	}

	return nil
}

func CheckResourceProviderOptions(options ResourceProviderOptions) error {
	err := options2.CheckWeb3Options(options.Web3)
	if err != nil {
		return err
	}
	err = CheckResourceProviderOfferOptions(options.Offers)
	if err != nil {
		return err
	}
	err = options2.CheckServicesOptions(options.Offers.Services)
	if err != nil {
		return err
	}
	err = options2.CheckBacalhauOptions(options.Bacalhau)
	if err != nil {
		return err
	}
	return nil
}

func ProcessResourceProviderOfferOptions(options ResourceProviderOfferOptions) (ResourceProviderOfferOptions, error) {
	// if there are no specs then populate with the single spec
	if len(options.Specs) == 0 {
		// loop the number of machines we want to offer
		for i := 0; i < options.OfferCount; i++ {
			options.Specs = append(options.Specs, options.OfferSpec)
		}
	}
	return options, nil
}

func ProcessResourceProviderOptions(options ResourceProviderOptions) (ResourceProviderOptions, error) {
	newOfferOptions, err := ProcessResourceProviderOfferOptions(options.Offers)
	if err != nil {
		return options, err
	}
	options.Offers = newOfferOptions
	newWeb3Options, err := options2.ProcessWeb3Options(options.Web3)
	if err != nil {
		return options, err
	}
	options.Web3 = newWeb3Options
	return options, CheckResourceProviderOptions(options)
}
