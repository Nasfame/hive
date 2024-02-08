package config

import (
	"os"
	"path"

	"github.com/CoopHive/hive/enums"
)

var appConfig = configMap[string]{
	enums.DEBUG: {
		desc:       "debug mode",
		defaultVal: "false",
	},
	enums.DEALER: {
		"Dealer name for resource provider",
		"std-autoaccept", // std-autoaccept
	},
	// 	enums.DEALER_PATH: {
	// 		"Dealer path for resource provider",
	// 		"./autoaccept.so", // TODO: set default to empty
	// 	},

	enums.APP_DIR: {
		"App Location Directory",
		"$HOME/coophive",
	},
	// derived
	// enums.APP_DATA_DIR: {
	// 	"App Data Location: typically used for storing github repos and results",
	// 	"$APP_DIR/data",
	// },
	//
	// enums.APP_PLUGIN_DIR: {
	// 	"Plugin Path: typically used for storing plugins",
	// 	"$APP_DIR/plugins",
	// },

	enums.COOPHIVE_CONTROLLER_ADDRESS: { // TODO: network dependent
		"Web3 Controller Address",
		"0xCCAaFD2AdD790788436f10e2C84585C46388b9aF",
	},
}

func init() {
	userDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	appConfig[enums.APP_DIR].defaultVal = path.Join(userDir, "coophive")
}
