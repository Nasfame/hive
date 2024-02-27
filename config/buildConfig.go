package config

import (
	"errors"
	"os"
	"path"
	"slices"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/CoopHive/hive/enums"
	"github.com/CoopHive/hive/utils"
)

var version string
var commitSha string

// //go:embed version.txt
// var version string TODO: another way to embed

// // go:embed buildDate.txt
// var buildDate string
// TODO:
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
	enums.VERSION: {
		desc:       "version",
		defaultVal: version,
	},
	enums.COMMIT_SHA: {
		desc:       "commit sha",
		defaultVal: commitSha,
	},

	enums.RELEASE_URL: {
		desc:       "release url",
		defaultVal: "https://github.com/CoopHive/hive/releases",
	},

	enums.APP_LOG_FILE_FORMAT: {
		"app log file format",
		"/coophive_%s.jsonl", // injects appdir into the format
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
	enums.BACALHAU_RESULTS_DIR: {
		"bacalhau results dir", // relative to app dir
		"bacalhau-results",
	},
	enums.BACALHAU_SPECS_DIR: {
		"bacalhau specs dir", // relative to app dir
		"bacalhau-specs",
	},
	enums.BACALHAU_JOBS_DIR: {
		"bacalhau jobs dir", // relative to app dir
		"bacalhau-specs",
	},
	enums.REPO_DIR: {
		"repos dir", // relative to app dir
		"repos",
	},
	enums.DOWNlOADS_DIR: {
		"downloads dir",
		"downloaded-files", // relative to app dir
	},
	enums.JOB_PRICE: {
		"job price", // is now hardcoded, but TODO:
		"2",
	},
}

// TODO: add network related contract configs but keep it open to recevie from env

var APP_DATA_DIR string // use inject Conf , temporarily using this global variable
var MODULE_PATH string  // temporary init for shortcuts pkg, use Conf injected

var STD_MODULE_FORMAT string
var Conf *viper.Viper

func tempInitForFx(conf *viper.Viper) {

	Conf = conf // set global var
	// log.Println("version", conf.GetString(enums.VERSION))

	MODULE_PATH = conf.GetString(enums.MODULE_PATH)
	APP_DATA_DIR = conf.GetString(enums.APP_DATA_DIR)
	STD_MODULE_FORMAT = conf.GetString(enums.STD_MODULE_FORMAT)

	// log.Println("app data dir", APP_DATA_DIR)

	if conf.GetString(enums.NETWORK) == deprecatedNetwork {
		X_COOPHIVE_USER_HEADER = oldUserHeader
		X_COOPHIVE_SIGNATURE_HEADER = oldSignatureHeader
	}
	if conf.GetBool(enums.DEBUG) {
		// zerolog.SetGlobalLevel(zerolog.DebugLevel)
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
		log.Debug().Msgf("debug mode enabled")
	}

	appDir := conf.GetString(enums.APP_DIR)

	conf.Set(enums.APP_PLUGIN_DIR, path.Join(appDir, "plugins"))

	// SetAppDir(conf, appDir) deprecated as services are directly taking the r

	// log.Fatal().Msg(conf.GetString(enums.WEB3_PRIVATE_KEY))

}

func SetAppDir(conf *viper.Viper, appDir string) {

	appDir, err := utils.EnsureDir(appDir)
	if err != nil {
		logrus.Fatalf("failed to create app dir: %s due to %v\n", appDir, err)
	}
	conf.Set(enums.APP_DIR, appDir)

	setPathConfig := func(pathKey string) {
		dirPath := conf.GetString(pathKey)
		newPath := path.Join(appDir, dirPath)

		// err := os.MkdirAll(newPath, 0755)
		newPath, err := utils.EnsureDir(newPath)
		if !errors.Is(err, os.ErrExist) && err != nil {
			log.Fatal().Err(err).Msgf("failed to create dir: %s", newPath)
			return
		}

		conf.Set(pathKey, newPath)
		log.Debug().Msgf("set %s=%s", pathKey, newPath)

		log.Debug().Msgf("created dir: %s", newPath)

	}

	// conf.Set(enums.BACALHAU_RESULTS_DIR, path.Join(appDir, conf.GetString(enums.BACALHAU_RESULTS_DIR)))
	// os.MkdirAll(conf.GetString(enums.BACALHAU_RESULTS_DIR), 0755)

	setPathConfig(enums.BACALHAU_RESULTS_DIR)
	setPathConfig(enums.BACALHAU_SPECS_DIR)
	setPathConfig(enums.BACALHAU_JOBS_DIR)
	setPathConfig(enums.REPO_DIR)
	setPathConfig(enums.DOWNlOADS_DIR)

	setPathConfig(enums.APP_LOG_FILE_FORMAT)

	bacalhauEnv := conf.GetString(enums.BACALHAU_ENV)

	bacalhauEnvs := strings.Split(bacalhauEnv, "\n")

	bacalhauEnvs = slices.Compact(bacalhauEnvs)

	conf.Set(enums.BACALHAU_ENV, bacalhauEnvs)

}
