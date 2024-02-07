package config

import (
	"github.com/CoopHive/hive/enums"
)

var appConfig = configMap[string]{
	enums.DEBUG: {
		desc:       "debug mode",
		defaultVal: "false",
	},
	enums.DEALER: {
		"Dealer name for resource provider",
		"autoaccept", // std-autoaccept
	},
	// 	enums.DEALER_PATH: {
	// 		"Dealer path for resource provider",
	// 		"./autoaccept.so", // TODO: set default to empty
	// 	},

	enums.APP_DIR: {
		"App Location Directory",
		"/tmp/coophive",
	},

	enums.APP_DATA_DIR: {
		"App Data Location: typically used for storing github repos and results",
		"/tmp/coophive/data",
	},

	enums.APP_PLUGIN_DIR: {
		"Plugin Path: typically used for storing plugins",
		"/tmp/coophive/plugins",
	},
}
