package config

import (
	"fmt"
	"os"
	"path"

	"github.com/CoopHive/hive/enums"
)

const DEFAULT_DEALER = "std-autoaccept"

const defaultNetwork = "sepolia"

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

	enums.HIVE_CONTROLLER: { // TODO: network dependent
		"Web3 Controller Address",
		"0xCCAaFD2AdD790788436f10e2C84585C46388b9aF",
	},
	enums.HIVE_SOLVER: {
		"Hive Solver  Address",
		"0xB4A1671063fe482a95C2519b78E3974EFAd87854",
	},
	enums.HIVE_MEDIATION: {
		"Hive Mediation Addresss : can be set of addresses separated by ','",
		"0x2d83ced7562e406151bd49c749654429907543b4",
	},
	enums.NETWORK: {
		fmt.Sprintf("supported networks:%v. aurora is deprecated", NETWORKS),
		defaultNetwork,
	},

	enums.WEB3_RPC_URL: {
		"rpc url",
		"ws://testnet.co-ophive.network:8546",
	},
	enums.WEB3_CHAIN_ID: {
		"chain id of the network",
		"1337",
	},
	enums.WEB3_PRIVATE_KEY: {
		"private key",
		"0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
	},
	enums.RP_PRIVATE_KEY: {
		"private key",
		"",
	},
	enums.JC_PRIVATE_KEY: {
		"private key",
		"",
	},
	enums.MEDIATOR_PRIVATE_KEY: {
		"private key",
		"",
	},
	enums.SOLVER_PRIVATE_KEY: {
		"private key",
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

	/*RP*/

	enums.BACALHAU_API_HOST: {
		"bacalhau host",
		"localhost",
	},

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

func init() {
	userDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	appConfig[enums.APP_DIR].defaultVal = path.Join(userDir, "coophive")
}
