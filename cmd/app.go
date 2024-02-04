package cmd

import (
	"log"

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

		// fx.StartTimeout(0),
		// fx.StopTimeout(1),
	)
	app.Run()
	<-app.Done()
	log.Println("exiting gracefully")

}
