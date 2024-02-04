package config

import (
	"log"
	"os"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"

	"github.com/CoopHive/hive/enums"
)

var Module = fx.Options(
	fx.Provide(
		newConfig,
	),
	fx.Invoke(tempInitForFx),
	fx.WithLogger(func(conf *viper.Viper) (l fxevent.Logger) {
		if conf.GetBool(enums.DEBUG) {
			return &fxevent.ConsoleLogger{W: os.Stderr}
		}
		return fxevent.NopLogger
	}))

type out struct {
	fx.Out

	Conf *viper.Viper
}

func newConfig() (o out) {
	pf := pflag.NewFlagSet("conf", pflag.ContinueOnError)
	pf.Parse(os.Args[1:])

	// fmt.Println(os.Args)

	config := viper.New()

	checkDup := func(key string, block string) {
		if config.IsSet(key) {
			log.Fatalf("duplicate key found in config[%s]: %s", block, key)
		}

	}

	for key, meta := range buildConfig {
		checkDup(key, "build")
		config.Set(key, meta.defaultVal)
	}

	for key, meta := range appConfig {
		checkDup(key, "app")

		config.SetDefault(key, meta.defaultVal)

		// automatic conversion of environment var key to `UPPER_CASE` will happen.
		config.BindEnv(key)
	}

	for key, meta := range jobCreatorConfig {
		checkDup(key, "jobCreator")

		config.SetDefault(key, meta.defaultVal)

		// automatic conversion of environment var key to `UPPER_CASE` will happen.
		config.BindEnv(key)

		if key == enums.APP_DATA_DIR {
			// read command-line arguments
			pf.String(key, meta.defaultVal, meta.desc)
			pflag.String(key, meta.defaultVal, meta.desc) // to show in usage
		}
	}

	config.BindPFlags(pf)

	o.Conf = config

	return
}

func debuggerOptions(conf *viper.Viper) fx.Option {

	if conf.GetBool(enums.DEBUG) {
		log.Fatal("DEBUG MODE")
		return fx.Options()
	}

	return fx.NopLogger

}
