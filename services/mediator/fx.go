package mediator

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

	MediatorCmd *cobra.Command `name:"mediator"`
}

func newServices(i in) (o out) {

	cmd := newMediatorCmd()

	o = out{
		MediatorCmd: cmd,
	}
	return
}
