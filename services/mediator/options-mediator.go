package mediator

import (
	"fmt"

	options2 "github.com/CoopHive/hive/pkg/options"
	"github.com/CoopHive/hive/pkg/system"
	"github.com/CoopHive/hive/services/internal-mediator/mediator"

	"github.com/spf13/cobra"
)

func NewMediatorOptions() mediator.MediatorOptions {
	options := mediator.MediatorOptions{
		Bacalhau: options2.GetDefaultBacalhauOptions(),
		Web3:     options2.GetDefaultWeb3Options(),
		Services: options2.GetDefaultServicesOptions(),
	}
	options.Web3.Service = system.MediatorService
	return options
}

func AddMediatorCliFlags(cmd *cobra.Command, options *mediator.MediatorOptions) {
	options2.AddBacalhauCliFlags(cmd, &options.Bacalhau)
	options2.AddWeb3CliFlags(cmd, &options.Web3)
	options2.AddServicesCliFlags(cmd, &options.Services)
}

func CheckMediatorOptions(options mediator.MediatorOptions) error {
	err := options2.CheckWeb3Options(options.Web3)
	if err != nil {
		return err
	}
	err = options2.CheckBacalhauOptions(options.Bacalhau)
	if err != nil {
		return err
	}
	// only check the solver because we are the mediator
	if options.Services.Solver == "" {
		return fmt.Errorf("No solver service specified - please use SERVICE_SOLVER or --service-solver")
	}
	return nil
}

func ProcessMediatorOptions(options mediator.MediatorOptions) (mediator.MediatorOptions, error) {
	newWeb3Options, err := options2.ProcessWeb3Options(options.Web3)
	if err != nil {
		return options, err
	}
	options.Web3 = newWeb3Options
	return options, CheckMediatorOptions(options)
}
