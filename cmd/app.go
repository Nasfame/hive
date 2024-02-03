package cmd

import (
	"go.uber.org/fx"

	"github.com/CoopHive/hive/config"
	"github.com/CoopHive/hive/internal"
	"github.com/CoopHive/hive/services"
)

func Hive() {

	app := fx.New(
		config.Module,
		internal.Module,
		services.Module,

		fx.StartTimeout(0),
		fx.StopTimeout(0),

		fx.NopLogger, // FIXME: move to config and get this configured async
	)
	app.Run()
}
