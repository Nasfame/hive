package resourceprovider

import (
	"github.com/spf13/cobra"
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
}

type out struct {
	fx.Out

	ResourceProviderCmd *cobra.Command `name:"internal-resourceprovider"`
}

func newServices(i in) (o out) {

	cmd := newResourceProviderCmd()

	o = out{
		ResourceProviderCmd: cmd,
	}
	return
}
