package solver

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

	SolverCmd *cobra.Command `name:"solver"`
}

func newServices(i in) (o out) {

	cmd := newSolverCmd()

	o = out{
		SolverCmd: cmd,
	}
	return
}
