package cmd

import (
	"log"
	"os"

	"go.uber.org/fx"

	"github.com/CoopHive/hive/config"
	"github.com/CoopHive/hive/internal"
	"github.com/CoopHive/hive/services"
)

func Hive() {

	app := fx.New(
		// fx.Provide(zap.NewProduction), // TODO:

		//
		// fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
		// 	if os.Getenv("DEBUG") == "true" {
		// 		log.Println("DEBUG MODE")
		// 		return &fxlog.DefaultLogger(os.Stderr)
		// 	}
		//
		// 	return fxevent.NopLogger
		// }),

		config.Module,
		internal.Module,
		services.Module,

		fx.StartTimeout(0),
		fx.StopTimeout(1),
		debuggerOptions(), // FIXME: after testing debuggerOptions is working or not in config/fx.go
	)
	app.Run()
	<-app.Done()
	log.Println("exiting gracefully")

}

func debuggerOptions() fx.Option {

	if os.Getenv("DEBUG") == "true" {
		log.Println("DEBUG MODE")
		return fx.Provide()
	}

	// log.Println("PRODUCTION=true")

	return fx.NopLogger

}
