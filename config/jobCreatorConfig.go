package config

import (
	"github.com/CoopHive/hive/enums"
)

var jobCreatorConfig = configMap[string]{
	enums.APP_DATA_DIR: {
		"app data directory : for storing github repos and downloaded results",
		"/tmp/coophive/data",
	},
}

var APP_DATA_DIR string // use inject Conf , temporarily using this global variable
