package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"github.com/CoopHive/hive/enums"
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
func init() {
	StdRepoUri := buildConfig[enums.STD_REPO_URI].defaultVal
	StdModulePrefix := buildConfig[enums.STD_MODULE_PREFIX].defaultVal
	CoophiveStdModule := StdRepoUri + StdModulePrefix + "-%s"

	stdModuleFormat := buildConfig[enums.STD_MODULE_FORMAT]
	stdModuleFormat.defaultVal = CoophiveStdModule

}

func init() {
	configFile := os.Getenv("CONFIG_FILE")

	if configFile == "" {
		configFile = ".env"
	}

	logrus.Debugf("Loading config from %s", configFile)

	if err := godotenv.Load(configFile); err != nil {
		logrus.Debugf(".env loading error %v", err)
	}

}
