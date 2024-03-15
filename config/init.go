package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/CoopHive/hive/enums"
	"github.com/CoopHive/hive/utils"
)

/*
	func init() {
		userDir, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}

		appConfig[enums.APP_DIR].defaultVal = path.Join(userDir, "coophive")
	}
*/

// set debug
func init() {
	debugFlag, _ := strconv.ParseBool(os.Getenv(enums.DEBUG))

	if debugFlag {
		logrus.SetLevel(logrus.DebugLevel)
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}

// set module defaults
func init() {
	StdRepoUri := buildConfig[enums.STD_REPO_URI].defaultVal
	StdModulePrefix := buildConfig[enums.STD_MODULE_PREFIX].defaultVal
	CoophiveStdModule := StdRepoUri + StdModulePrefix + "-%s"

	stdModuleFormat := buildConfig[enums.STD_MODULE_FORMAT]
	stdModuleFormat.defaultVal = CoophiveStdModule

}

// load ConfigFile
func init() {
	configFile := os.Getenv("CONFIG_FILE")
	defaultLoad := false

	if configFile == "" {
		configFile = ".env"
		defaultLoad = true
	}

	logrus.Debugf("Loading config from %s", configFile)

	if err := godotenv.Load(configFile); err != nil {
		if !defaultLoad {
			logrus.Errorf(".env loading error %v", err)
		} else {
			logrus.Debugf("err loading : %v", err)
		}
	}

}

// temp init basics for tests

func init() {
	config := viper.New() // overriden by fx

	checkDup := func(key string, block string) {
		if config.IsSet(key) {
			err := fmt.Errorf("duplicate key found in config[%s]: %s", block, key)
			panic(err)
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
		if err := config.BindEnv(key); err != nil {
			panic(err)
		}
	}

	for keyArg, meta := range featureFlags {
		key := keyArg.String()

		checkDup(key, "featureFlag")

		config.SetDefault(key, fmt.Sprint(meta.defaultVal))

		envValue := os.Getenv(key)
		if envValue != "" {
			// TODO: check why that is not working
			config.Set(key, envValue)
		}
	}

	Conf = config // overriden by fx TODO: perhaps migrate this

	initDerivedConfigVariables(config)

	utils.EnsureDir(config.GetString(enums.BACALHAU_SERVE_IPFS_PATH))

}
