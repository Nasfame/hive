package jobcreator

import (
	"github.com/spf13/cobra"
	"go.uber.org/fx"

	"github.com/CoopHive/hive/internal/genesis"
)

var Module = fx.Options(
	fx.Provide(
		newServices,
	),
	// fx.Supply() //TODO: initialize a job creator for run command
)

type in struct {
	fx.In
	*genesis.Service
}

type out struct {
	fx.Out

	JobCreatorCmd *cobra.Command `name:"jobcreator"`
}

func newServices(i in) (o out) {

	cmd := newJobCreatorCmd()

	o = out{
		JobCreatorCmd: cmd,
	}
	return
}
