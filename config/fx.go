package config

import (
	"log"

	"github.com/spf13/viper"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		newConfig,
	),
	fx.Invoke(tempInitForFx),
	fx.Provide(debuggerOptions),
)

func newConfig() (config *viper.Viper) {
	config = viper.New()

	checkDup := func(key string, block string) {
		if config.IsSet(key) {
			log.Fatalf("duplicate key found in config[%s]: %s", block, key)
		}

	}

	for key, val := range buildConfig {
		checkDup(key, "build")
		config.Set(key, val.defaultVal)

	}

	for key, _ := range jobCreatorConfig {
		checkDup(key, "jobCreator")

		// automatic conversion of environment var key to `UPPER_CASE` will happen.
		config.BindEnv(key)
	}

	// viper.AutomaticEnv() //replaces the default values//not required really
	// pflag.Parse() //TODO: try to take
	// config.BindPFlags(pflag.CommandLine)
	return
}

func debuggerOptions(conf *viper.Viper) fx.Option {

	if conf.GetBool("DEBUG") {
		log.Println("DEBUG MODE")
		return fx.Options()
	}

	return fx.NopLogger

}
