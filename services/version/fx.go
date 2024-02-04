package version

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

	VersionCmd *cobra.Command `name:"version"` // TODO: use versionCmd tag
}

func newServices(i in) (o out) {

	cmd := newVersionCmd(i.Conf)

	o = out{
		VersionCmd: cmd,
	}
	return
}
