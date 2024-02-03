package jobcreator

import (
	"github.com/spf13/cobra"

	"github.com/CoopHive/hive/pkg/jobcreator"
	optionsfactory "github.com/CoopHive/hive/pkg/options"
	"github.com/CoopHive/hive/pkg/system"
	"github.com/CoopHive/hive/pkg/web3"
)

func newJobCreatorCmd() *cobra.Command {
	options := optionsfactory.NewJobCreatorOptions()

	cmd := &cobra.Command{
		Use:     "jobcreator",
		Short:   "Start the CoopHive job creator service.",
		Long:    "Start the CoopHive job creator service.",
		Example: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			options, err := optionsfactory.ProcessOnChainJobCreatorOptions(options, args)
			if err != nil {
				return err
			}
			return runJobCreator(cmd, options)
		},
	}

	optionsfactory.AddJobCreatorCliFlags(cmd, &options)

	return cmd
}

func runJobCreator(cmd *cobra.Command, options jobcreator.JobCreatorOptions) error {
	commandCtx := system.NewCommandContext(cmd)
	defer commandCtx.Cleanup()

	web3SDK, err := web3.NewContractSDK(options.Web3)
	if err != nil {
		return err
	}

	// create the job creator and start it's control loop
	jobCreatorService, err := jobcreator.NewOnChainJobCreator(options, web3SDK)
	if err != nil {
		return err
	}

	jobCreatorErrors := jobCreatorService.Start(commandCtx.Ctx, commandCtx.Cm)

	for {
		select {
		case err := <-jobCreatorErrors:
			return err
		case <-commandCtx.Ctx.Done():
			return nil
		}
	}
}
