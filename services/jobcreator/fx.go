package jobcreator

import (
	"github.com/spf13/cobra"
	"go.uber.org/fx"

	"github.com/CoopHive/hive/config"
	"github.com/CoopHive/hive/internal/genesis"
	"github.com/CoopHive/hive/internal/jobCreatorService"
	"github.com/CoopHive/hive/pkg/system"
	"github.com/CoopHive/hive/pkg/web3"
	"github.com/CoopHive/hive/services/dealmaker"
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

	DealMakerService *dealmaker.Service `name:"dealmaker"`
}

type out struct {
	fx.Out

	JobCreatorCmd *cobra.Command `name:"jc"`
}

func newServices(i in) (o out) {

	s := service{
		i.DealMakerService,
		i.Service,
	}

	cmd := s.newJobCreatorCmd()

	o = out{
		JobCreatorCmd: cmd,
	}
	return
}

func (s *service) newJobCreatorCmd() *cobra.Command {
	options := NewJobCreatorOptions()

	cmd := &cobra.Command{
		Use:     "jobcreator",
		Aliases: []string{"jc", "job-creator"},
		Short:   "Start the CoopHive job creator service.",
		Long:    "Start the CoopHive job creator service.",
		Example: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			options, err := ProcessOnChainJobCreatorOptions(options, args)
			if err != nil {
				return err
			}
			return s.runJobCreator(cmd, options)
		},
	}

	AddJobCreatorCliFlags(cmd, &options)

	return cmd
}

func (s *service) runJobCreator(cmd *cobra.Command, options jobCreatorService.JobCreatorOptions) error {
	commandCtx := system.NewCommandContext(cmd)
	defer commandCtx.Cleanup()

	if options.Dealer != config.DEFAULT_DEALER {
		if err := s.dealMakerService.LoadPlugin(options.Dealer); err != nil {
			s.Log.Errorf("Dealer %s is not supported on this platform", options.Dealer)
		}
	}

	web3SDK, err := web3.NewContractSDK(options.Web3)
	if err != nil {
		return err
	}

	// create the job creator and start it's control loop
	jobCreatorService, err := jobCreatorService.NewOnChainJobCreator(options, web3SDK, s.dealMakerService)
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
