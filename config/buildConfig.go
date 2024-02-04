package config

import (
	"github.com/spf13/viper"

	"github.com/CoopHive/hive/enums"
)

var version string
var commit_sha string

var buildConfig = configMap[string]{
	// app specific
	enums.APP_NAME: {
		"app name",
		"CoopHive",
	},
	enums.ENV: {
		"environment",
		enums.DEV,
	},
	enums.DEBUG: {
		desc:       "debug mode",
		defaultVal: "false",
	},

	enums.VERSION: {
		desc:       "version",
		defaultVal: version,
	},
	enums.COMMIT_SHA: {
		desc:       "commit sha",
		defaultVal: commit_sha,
	},

	enums.RELEASE_URL: {
		desc:       "release url",
		defaultVal: "https://github.com/CoopHive/hive/releases",
	},

	enums.APP_LOG_FILE_FORMAT: {
		"app log file format",
		"/var/tmp/coophive_%s.jsonl",
	},

	enums.STD_REPO_URI: {
		"coophive std module hosted uri",
		"https://github.com/CoopHive/",
	},

	enums.STD_MODULE_PREFIX: {
		"coophive std module prefix",
		"coophive-module",
	},

	enums.STD_MODULE_FORMAT: {
		"coophive std module format",
		"", // to be init
	},

	enums.MODULE_PATH: {
		"module path",
		"/module.coophive",
	},
}

var MODULE_PATH string // temporary init for shortcuts pkg, use conf injected

var STD_MODULE_FORMAT string
var Conf *viper.Viper

func init() {
	StdRepoUri := buildConfig[enums.STD_REPO_URI].defaultVal
	StdModulePrefix := buildConfig[enums.STD_MODULE_PREFIX].defaultVal
	CoophiveStdModule := StdRepoUri + StdModulePrefix + "-%s"

	stdModuleFormat := buildConfig[enums.STD_MODULE_FORMAT]
	stdModuleFormat.defaultVal = CoophiveStdModule

}

func tempInitForFx(conf *viper.Viper) {

	// StdRepoUri := buildConfig[enums.STD_REPO_URI].defaultVal
	// StdModulePrefix := buildConfig[enums.STD_MODULE_PREFIX].defaultVal
	// CoophiveStdModule := StdRepoUri + StdModulePrefix + "-%s"
	//
	// stdModuleFormat := buildConfig[enums.STD_MODULE_FORMAT]
	// stdModuleFormat.defaultVal = CoophiveStdModule

	// conf.Set(enums.STD_MODULE_FORMAT, stdModuleFormat.defaultVal)

	Conf = conf

	MODULE_PATH = conf.GetString(enums.MODULE_PATH)
	AppDataDir = conf.GetString(enums.APP_DATA_DIR)

	STD_MODULE_FORMAT = conf.GetString(enums.STD_MODULE_FORMAT)

	// log.Println("app dat dir", AppDataDir)

}
