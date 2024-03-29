package mediator

import (
	"github.com/spf13/cobra"
	"go.uber.org/fx"

	"github.com/CoopHive/hive/config"
	"github.com/CoopHive/hive/enums"
	"github.com/CoopHive/hive/internal/genesis"
)

var Module = fx.Options(
	fx.Provide(
		newServices,
	),
)

type in struct {
	fx.In
	*genesis.Service
}

type out struct {
	fx.Out

	MediatorCmd *cobra.Command `name:"mediator"`
}

func newServices(i in) (o out) {

	cmd := newMediatorCmd()

	o = out{
		MediatorCmd: cmd,
	}
	return
}

func newMediatorCmd() *cobra.Command {
	options := NewMediatorOptions()

	serviceType := enums.MEDIATOR

	mediatorCmd := &cobra.Command{
		Use:     "mediator",
		Aliases: []string{"mediate"},
		Short:   "Start the CoopHive mediator service.",
		Long:    "Start the CoopHive mediator service.",
		Example: "",
		RunE: func(cmd *cobra.Command, _ []string) error {
			serviceType.ProcessSyncServiceDirectory(config.Conf.GetString(enums.APP_DIR), func(appDir string) {
				config.SetAppDir(config.Conf, appDir)
			})
			options, err := ProcessMediatorOptions(options)
			if err != nil {
				return err
			}
			return runMediator(cmd, options)
		},
	}

	AddMediatorCliFlags(mediatorCmd, &options)

	return mediatorCmd
}
