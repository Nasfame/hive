package config

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/CoopHive/hive/enums"
)

const DEFAULT_DEALER = "std-autoaccept"

const defaultNetwork = "aurora"

const deprecatedNetwork = "aurora"

const AppDirSymbol = "$APP_DIR"

var processAppDir = func() string {

	appDir := buildConfig[enums.APP_NAME].defaultVal

	if strings.TrimSpace(appDir) == "" {
		log.Fatal("invalid app dir")
	}

	appDir = "." + strings.ToLower(appDir) // capital dirs are treated differently in unix systems

	homeDir, err := os.UserHomeDir()

	if err != nil {
		log.Fatal("failed to find user home dir with error ", err)
	}

	return path.Join(homeDir, appDir)
}

var appConfig = configMap[string]{
	enums.DEBUG: {
		desc:       "debug mode",
		defaultVal: "false",
	},
	enums.DEALER: {
		"Dealer for accepting/denying solver-matched deals",
		DEFAULT_DEALER,
	},
	// 	enums.DEALER_PATH: {
	// 		"Dealer path for resource provider",
	// 		"./autoaccept.so", // TODO: set default to empty
	// 	},

	enums.APP_DIR: {
		"App Location Directory",
		processAppDir(),
	},
	// derived
	enums.APP_DATA_DIR: { // no longer derived
		"App Data Location: typically used for storing github repos and results",
		AppDirSymbol + "/data",
	},
	//
	// enums.APP_PLUGIN_DIR: {
	// 	"Plugin Path: typically used for storing plugins",
	// 	"$APP_DIR/plugins",
	// },

	enums.HIVE_CONTROLLER: { // TODO: network dependent
		"Web3 Controller Address",
		"",
	},
	enums.HIVE_SOLVER: {
		"Hive Solver  Address",
		"",
	},
	enums.HIVE_MEDIATORS: {
		"Hive Mediation Addresss : can be set of addresses separated by ','",
		"",
	},
	enums.NETWORK: {
		fmt.Sprintf("supported networks:%v. aurora is deprecated", NETWORKS),
		defaultNetwork,
	},

	enums.WEB3_RPC_URL: {
		"rpc url",
		"",
	},
	enums.WEB3_CHAIN_ID: {
		"chain id of the network",
		"",
	},
	enums.WEB3_PRIVATE_KEY: {
		"private key",
		"",
		// "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
	},
	enums.HIVE_PRIVATE_KEY: {
		"hive private key: exclusive overrides web3 private key",
		"",
	},
	enums.RP_PRIVATE_KEY: {
		"private key for rp (overrides)",
		"",
	},
	enums.JC_PRIVATE_KEY: {
		"private key for jc",
		"",
	},
	enums.MEDIATOR_PRIVATE_KEY: {
		"private key for mediator",
		"",
	},
	enums.SOLVER_PRIVATE_KEY: {
		"private key for solver",
		"",
	},
	enums.SERVER_HOST: {
		"server host",
		"0.0.0.0",
	},

	enums.SERVER_PORT: {
		"server port",
		"8080",
	},
	enums.SERVER_URL: {
		"public facing url without application protocol like tcp, http",
		"",
	},

	/*Bacalhau*/

	enums.BACALHAU_BIN: {
		"bacalhau binary path: if its installed then its just bacalhau",
		"bacalhau",
	},

	enums.BACALHAU_API_HOST: {
		"bacalhau host",
		"localhost",
	},
	enums.BACALHAU_HOME: {
		"bacalhau home",
		"",
	},
	enums.BACALHAU_ENV: {
		"bacalhau env", // bacalhau env
		"",
	},

	/*RP*/

	enums.PRICING_MODE: {
		"pricing mode in integer",
		"",
		// 	https://github.com/CoopHive/hive/blob/8ce2279f1a77f60f933f53e04569878115749165/pkg/options/options-pricing.go#L14
	},

	enums.OFFER_CPU: {
		`How many milli-cpus to offer the network (OFFER_CPU).`,
		"1000",
	},

	enums.OFFER_GPU: {
		`How many milli-gpus to offer the network (OFFER_GPU).`,
		"0",
	},

	enums.OFFER_RAM: {
		`How many megabytes of RAM to offer the network (OFFER_RAM).`,
		"1024",
	},

	enums.OFFER_COUNT: {
		`How many machines will we offer using the cpu, ram and gpu settings (OFFER_COUNT).`,
		"1",
	},

	enums.OFFER_MODULES: {
		`The modules you are willing to run (OFFER_MODULES). Provide in , separated string`,
		"",
	},

	/*JC*/
	enums.MODULE_NAME: {
		"module name",
		"",
	},
	enums.MODULE_REPO: {
		"module repo",
		"",
	},
	enums.MODULE_HASH: {
		"module path",
		"",
	},
}
