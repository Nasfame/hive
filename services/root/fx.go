package root

import (
	"context"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/fx"

	"github.com/CoopHive/hive/internal/genesis"
)

var Module = fx.Options(
	fx.Provide(
		newServices,
	),
	fx.Invoke(executeRunCmd),
)

type in struct {
	fx.In
	*genesis.Service

	Conf *viper.Viper

	VersionCmd          *cobra.Command `name:"version"`
	RunCmd              *cobra.Command `name:"run"`
	JobCreatorCmd       *cobra.Command `name:"jobcreator"`
	ResourceProviderCmd *cobra.Command `name:"internal-resourceprovider"`
	MediatorCmd         *cobra.Command `name:"mediator"`
	SolverCmd           *cobra.Command `name:"solver"`
}

type out struct {
	fx.Out

	RootCmd *cobra.Command `name:"root""`
}

func newServices(i in) (o out) {
	cmd := newRootCmd(i.Conf, i.VersionCmd, i.RunCmd, i.JobCreatorCmd, i.ResourceProviderCmd, i.MediatorCmd, i.SolverCmd)

	o = out{
		RootCmd: cmd,
	}

	return
}

type inExec struct {
	fx.In

	RootCmd *cobra.Command `name:"root"`
}

func executeRunCmd(i inExec) {
	cmd := i.RootCmd

	cmd.SetContext(context.Background())
	cmd.SetOutput(os.Stdout)
	if err := cmd.Execute(); err != nil {
		Fatal(cmd, err.Error(), 1)
	}

}
