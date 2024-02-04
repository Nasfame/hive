package jobcreator

import (
	"github.com/spf13/cobra"

	"github.com/CoopHive/hive/pkg/system"
	"github.com/CoopHive/hive/pkg/web3"
	"github.com/CoopHive/hive/services/jobcreator/internal-job"
)

func newJobCreatorCmd() *cobra.Command {
	options := NewJobCreatorOptions()

	cmd := &cobra.Command{
		Use:     "jobcreator",
		Short:   "Start the CoopHive job creator service.",
		Long:    "Start the CoopHive job creator service.",
		Example: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			options, err := ProcessOnChainJobCreatorOptions(options, args)
			if err != nil {
				return err
			}
			return runJobCreator(cmd, options)
		},
	}

	AddJobCreatorCliFlags(cmd, &options)

	return cmd
}

func runJobCreator(cmd *cobra.Command, options internal_job.JobCreatorOptions) error {
	commandCtx := system.NewCommandContext(cmd)
	defer commandCtx.Cleanup()

	web3SDK, err := web3.NewContractSDK(options.Web3)
	if err != nil {
		return err
	}

	// create the job creator and start it's control loop
	jobCreatorService, err := internal_job.NewOnChainJobCreator(options, web3SDK)
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
