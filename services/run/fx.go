package run

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/fx"

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
	Conf *viper.Viper
}

type out struct {
	fx.Out

	RunCmd *cobra.Command `name:"run"`
}

func newServices(i in) (o out) {

	cmd := newRunCmd(i.Conf)

	o = out{
		RunCmd: cmd,
	}
	return
}