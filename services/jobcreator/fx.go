package jobcreator

import (
	"github.com/spf13/cobra"
	"go.uber.org/fx"

	"github.com/CoopHive/hive/internal/genesis"
	"github.com/CoopHive/hive/internal/jobCreatorService"
	"github.com/CoopHive/hive/pkg/system"
	"github.com/CoopHive/hive/pkg/web3"
)

var Module = fx.Options(
	fx.Provide(
		newServices,
	),
	// fx.Supply() //TODO: initialize a job creator for run command
)

type in struct {
	fx.In
	*genesis.Service
}

type out struct {
	fx.Out

	JobCreatorCmd *cobra.Command `name:"jobcreator"`
}

func newServices(i in) (o out) {

	cmd := newJobCreatorCmd()

	o = out{
		JobCreatorCmd: cmd,
	}
	return
}

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

func runJobCreator(cmd *cobra.Command, options jobCreatorService.JobCreatorOptions) error {
	commandCtx := system.NewCommandContext(cmd)
	defer commandCtx.Cleanup()

	web3SDK, err := web3.NewContractSDK(options.Web3)
	if err != nil {
		return err
	}

	// create the job creator and start it's control loop
	jobCreatorService, err := jobCreatorService.NewOnChainJobCreator(options, web3SDK)
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
