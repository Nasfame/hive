package mediator

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/CoopHive/hive/pkg/executor/bacalhau"
	"github.com/CoopHive/hive/pkg/mediator"
	optionsfactory "github.com/CoopHive/hive/pkg/options"
	"github.com/CoopHive/hive/pkg/system"
	"github.com/CoopHive/hive/pkg/web3"
)

func newMediatorCmd() *cobra.Command {
	options := optionsfactory.NewMediatorOptions()

	mediatorCmd := &cobra.Command{
		Use:     "mediator",
		Short:   "Start the CoopHive mediator service.",
		Long:    "Start the CoopHive mediator service.",
		Example: "",
		RunE: func(cmd *cobra.Command, _ []string) error {
			options, err := optionsfactory.ProcessMediatorOptions(options)
			if err != nil {
				return err
			}
			return runMediator(cmd, options)
		},
	}

	optionsfactory.AddMediatorCliFlags(mediatorCmd, &options)

	return mediatorCmd
}

func runMediator(cmd *cobra.Command, options mediator.MediatorOptions) error {
	commandCtx := system.NewCommandContext(cmd)
	defer commandCtx.Cleanup()

	web3SDK, err := web3.NewContractSDK(options.Web3)
	if err != nil {
		return err
	}

	executor, err := bacalhau.NewBacalhauExecutor(options.Bacalhau)
	if err != nil {
		return err
	}

	mediatorService, err := mediator.NewMediator(options, web3SDK, executor)
	if err != nil {
		return err
	}

	log.Debug().Msgf("Starting mediator service.")
	mediatorErrors := mediatorService.Start(commandCtx.Ctx, commandCtx.Cm)
	for {
		select {
		case err := <-mediatorErrors:
			commandCtx.Cleanup()
			return err
		case <-commandCtx.Ctx.Done():
			return nil
		}
	}
}
