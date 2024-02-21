package root

import (
	"context"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/fx"

	"github.com/CoopHive/hive/internal/genesis"
)

var Module = fx.Options(
	fx.Provide(
		newServices,
	),
	fx.Invoke(executeCobraCommand),
)

type in struct {
	fx.In
	*genesis.Service

	Conf *viper.Viper

	VersionCmd          *cobra.Command `name:"version"`
	RunCmd              *cobra.Command `name:"run"`
	JobCreatorCmd       *cobra.Command `name:"jc"`
	ResourceProviderCmd *cobra.Command `name:"rp"`
	MediatorCmd         *cobra.Command `name:"mediator"`
	SolverCmd           *cobra.Command `name:"solver"`
}

type out struct {
	fx.Out

	RootCmd *cobra.Command `name:"root"`
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

func executeCobraCommand(i inExec) {
	cmd := i.RootCmd

	cmd.SetContext(context.Background())
	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)

	go func() {
		if err := cmd.Execute(); err != nil {
			log.Debug().Err(err).Caller(1).Msg("error running command")
			Fatal(cmd, err.Error(), 1)
		}

		os.Exit(0)
	}()

}
