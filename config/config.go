package config

import (
	"github.com/joho/godotenv"

	"github.com/CoopHive/hive/enums"
)

var LogFileFormat = "/var/tmp/coophive_%s.jsonl"

const STD_REPO_URI = "https://github.com/CoopHive/"

const STD_MODULE_PREFIX = "coophive-module"

const COOPHIVE_STD_MODULE = STD_REPO_URI + STD_MODULE_PREFIX + "-%s"

type argvMeta struct {
	desc       string
	defaultVal string
}

var confList = map[string]argvMeta{
	enums.APP_NAME: {
		"app name",
		"CoopHive",
	},
	enums.ENV: {
		"environment",
		enums.DEV,
	},
}

var VERSION string

var version string // do ld -w here and link to viper

var COMMIT_SHA string

const GO_BINARY_URL = "https://github.com/CoopHive/hive/releases/"

func init() {
	// fmt.Printf("CoopHive: %s\n", hive.VERSION)

	err := godotenv.Load()
	if err != nil {
		// log.Debug().Str("err", err.Error()).Msgf(".env not found")
		// TODO: Doesn't look good, add custom flag
	}

}
