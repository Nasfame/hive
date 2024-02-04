package config

import (
	"log"
	"os"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/fx"

	"github.com/CoopHive/hive/enums"
)

var Module = fx.Options(
	fx.Provide(
		newConfig,
	),
	fx.Invoke(tempInitForFx),
	fx.Provide(debuggerOptions),
)

func newConfig() (config *viper.Viper) {

	pf := pflag.NewFlagSet("conf", pflag.ContinueOnError)
	pf.Parse(os.Args[1:])

	// fmt.Println(os.Args)

	config = viper.New()

	checkDup := func(key string, block string) {
		if config.IsSet(key) {
			log.Fatalf("duplicate key found in config[%s]: %s", block, key)
		}

	}

	for key, meta := range buildConfig {
		checkDup(key, "build")
		config.Set(key, meta.defaultVal)

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

	return
}

func debuggerOptions(conf *viper.Viper) fx.Option {

	if conf.GetBool("DEBUG") {
		log.Println("DEBUG MODE")
		return fx.Options()
	}

	return fx.NopLogger

}
