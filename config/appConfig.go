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
}
