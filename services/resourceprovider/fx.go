package resourceprovider

import (
	"github.com/spf13/cobra"
	"go.uber.org/fx"

	"github.com/CoopHive/hive/internal/genesis"
	"github.com/CoopHive/hive/services/dealmaker"
)

var Module = fx.Options(
	dealmaker.Module,
	fx.Provide(
		newServices,
	),
)

type in struct {
	fx.In
	*genesis.Service

	DealMakerService *dealmaker.Service `name:"dealmaker"`
}

type out struct {
	fx.Out

	ResourceProviderCmd *cobra.Command `name:"internal-resourceprovider"`
}

func newServices(i in) (o out) {

	s := &service{
		i.DealMakerService,
		i.Service,
	}

	cmd := s.newResourceProviderCmd()

	o = out{
		ResourceProviderCmd: cmd,
	}
	return
}
