package run

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/fx"

	"github.com/CoopHive/hive/internal/genesis"
	"github.com/CoopHive/hive/services/dealmaker"

	optionsfactory "github.com/CoopHive/hive/services/jobcreator"
)

var Module = fx.Options(
	fx.Provide(
		newServices,
	),
)

type in struct {
	fx.In
	*genesis.Service
	Conf *viper.Viper

	DealMakerService *dealmaker.Service `name:"dealmaker"`
}

type out struct {
	fx.Out

	RunCmd *cobra.Command `name:"run"`
}

func newServices(i in) (o out) {

	s := service{
		i.DealMakerService,
		i.Service,
	}

	cmd := s.newRunCmd(i.Conf)

	o = out{
		RunCmd: cmd,
	}
	return
}

func (s *service) newRunCmd(conf *viper.Viper) *cobra.Command {
	options := optionsfactory.NewJobCreatorOptions()
	runCmd := &cobra.Command{
		Use:     "run",
		Short:   "Run a job on the CoopHive network.",
		Long:    "Run a job on the CoopHive network.",
		Example: "run cowsay:v0.0.1 -i Message=CoopHive",
		RunE: func(cmd *cobra.Command, args []string) error {
			options, err := optionsfactory.ProcessJobCreatorOptions(options, args)
			if err != nil {
				return err
			}
			return s.runJob(cmd, options, conf)
		},
	}

	optionsfactory.AddJobCreatorCliFlags(runCmd, &options)

	return runCmd
}
