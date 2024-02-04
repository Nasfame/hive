package config

import (
	"github.com/CoopHive/hive/enums"
)

var jobCreatorConfig = configMap[string]{
	enums.APP_DATA_DIR: {
		"App Data Location: typically used for storing github repos and results",
		"/tmp/coophive/data",
	},
	// enums.JOBCREATOR_OPTION_INPUT: {
	// 	"i : ",
	// 	"",
	// },
}

var AppDataDir string // use inject Conf , temporarily using this global variable
