package jobcreator

import (
	"fmt"

	"github.com/CoopHive/hive/config"
	"github.com/CoopHive/hive/enums"
	"github.com/CoopHive/hive/internal/jobCreatorService"
	"github.com/CoopHive/hive/pkg/dto"
	options2 "github.com/CoopHive/hive/pkg/options"
	"github.com/CoopHive/hive/pkg/system"

	"github.com/spf13/cobra"
)

func NewJobCreatorOptions() jobCreatorService.JobCreatorOptions {
	options := jobCreatorService.JobCreatorOptions{
		GetDefaultJobCreatorMediationOptions(),
		GetDefaultJobCreatorOfferOptions(),
		options2.GetDefaultWeb3Options(enums.JC),
		config.DEFAULT_DEALER,
	}
	options.Web3.Service = system.JobCreatorService
	return options
}

func GetDefaultJobCreatorMediationOptions() jobCreatorService.JobCreatorMediationOptions {
	return jobCreatorService.JobCreatorMediationOptions{
		CheckResultsPercentage: options2.GetDefaultServeOptionInt("MEDIATION_CHANCE", 0),
	}
}

func GetDefaultJobCreatorOfferOptions() jobCreatorService.JobCreatorOfferOptions {
	return jobCreatorService.JobCreatorOfferOptions{
		Module: GetDefaultModuleOptions(),
		// this is the default pricing mode for an JC
		Mode:     options2.GetDefaultPricingMode(dto.MarketPrice),
		Pricing:  options2.GetDefaultPricingOptions(),
		Timeouts: options2.GetDefaultTimeoutOptions(),
		Inputs:   map[string]string{},
		Services: options2.GetDefaultServicesOptions(),
	}
}

func AddJobCreatorMediationCliFlags(cmd *cobra.Command, mediationOptions *jobCreatorService.JobCreatorMediationOptions) {
	cmd.PersistentFlags().IntVar(
		&mediationOptions.CheckResultsPercentage,
		"mediation-chance",
		mediationOptions.CheckResultsPercentage,
		"The percentage chance we will check results",
	)
}

func AddJobCreatorOfferCliFlags(cmd *cobra.Command, offerOptions *jobCreatorService.JobCreatorOfferOptions) {
	// add the inputs that we will merge into the module template file
	cmd.PersistentFlags().StringToStringVarP(&offerOptions.Inputs, "input", "i", offerOptions.Inputs, "Input key-value pairs")

	options2.AddPricingModeCliFlags(cmd, &offerOptions.Mode)
	options2.AddPricingCliFlags(cmd, &offerOptions.Pricing)
	options2.AddTimeoutCliFlags(cmd, &offerOptions.Timeouts)
	AddModuleCliFlags(cmd, &offerOptions.Module)
	options2.AddServicesCliFlags(cmd, &offerOptions.Services)
}

func AddJobCreatorCliFlags(cmd *cobra.Command, options *jobCreatorService.JobCreatorOptions) {
	AddJobCreatorMediationCliFlags(cmd, &options.Mediation)
	options2.AddWeb3CliFlags(cmd, &options.Web3)
	AddJobCreatorOfferCliFlags(cmd, &options.Offer)
	cmd.PersistentFlags().StringVar(&options.Dealer, "dealer", options.Dealer, "Choose Dealer for accepting/denying solver-matched deals")
}

func CheckJobCreatorOptions(options jobCreatorService.JobCreatorOptions) error {
	err := options2.CheckWeb3Options(options.Web3)
	if err != nil {
		return err
	}
	err = CheckModuleOptions(options.Offer.Module)
	if err != nil {
		return err
	}
	err = options2.CheckServicesOptions(options.Offer.Services)
	if err != nil {
		return err
	}

	if options.Mediation.CheckResultsPercentage < 0 || options.Mediation.CheckResultsPercentage > 100 {
		return fmt.Errorf("mediation-chance must be between 0 and 100")
	}

	return nil
}

func ProcessJobCreatorOptions(options jobCreatorService.JobCreatorOptions, args []string) (jobCreatorService.JobCreatorOptions, error) {
	name := ""
	if len(args) == 1 {
		name = args[0]
	}

	if name != "" {
		options.Offer.Module.Name = name
	}

	moduleOptions, err := ProcessModuleOptions(options.Offer.Module)
	if err != nil {
		return options, err
	}
	options.Offer.Module = moduleOptions
	newWeb3Options, err := options2.ProcessWeb3Options(options.Web3)
	if err != nil {
		return options, err
	}
	options.Web3 = newWeb3Options
	return options, CheckJobCreatorOptions(options)
}

func ProcessOnChainJobCreatorOptions(options jobCreatorService.JobCreatorOptions, args []string) (jobCreatorService.JobCreatorOptions, error) {
	newWeb3Options, err := options2.ProcessWeb3Options(options.Web3)
	if err != nil {
		return options, err
	}
	options.Web3 = newWeb3Options

	err = options2.CheckWeb3Options(options.Web3)
	if err != nil {
		return options, err
	}
	err = options2.CheckServicesOptions(options.Offer.Services)
	if err != nil {
		return options, err
	}

	options.Mediation.CheckResultsPercentage = 0

	return options, nil
}
