package config

import (
	"github.com/CoopHive/hive/enums"
)

var appConfig = configMap[string]{
	enums.DEBUG: {
		desc:       "debug mode",
		defaultVal: "false",
	},
}
