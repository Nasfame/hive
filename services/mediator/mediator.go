package mediator

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/CoopHive/hive/pkg/executor/bacalhau"
	"github.com/CoopHive/hive/pkg/system"
	"github.com/CoopHive/hive/pkg/web3"
)

func runMediator(cmd *cobra.Command, options MediatorOptions) error {
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

	mediatorService, err := NewMediator(options, web3SDK, executor)
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
