package enums

const (
	APP_NAME = "APP_NAME"
	ENV      = "ENV"
	PORT     = "PORT"
	TIMEZONE = "TIMEZONE"
)

const (
	MODE = "MODE"
)

const (
	DEV        = "DEV"
	PRODUCTION = "PRODUCTION"
)

const (
	DEBUG      = "DEBUG"
	VERSION    = "VERSION"
	COMMIT_SHA = "COMMIT_SHA"

	RELEASE_URL = "RELEASE_URL"
)

const (
	APP_LOG_FILE_FORMAT = "APP_LOG_FILE_FORMAT"
	STD_REPO_URI        = "STD_REPO_URI"

	STD_MODULE_PREFIX = "STD_MODULE_PREFIX"

	STD_MODULE_FORMAT = STD_REPO_URI + STD_MODULE_PREFIX + "-%s"
)

const (
	APP_DIR        = "app_dir"
	APP_DATA_DIR   = "app-data-dir"
	APP_PLUGIN_DIR = "app-plugin-dir"
)

const (
	JOBCREATOR_OPTION_INPUT = "i"
)

const (
	DEALER      = "dealer"
	DEALER_PATH = "dealer-path"
)

const (
	HIVE_CONTROLLER = "hive_controller"
	HIVE_PAYMENTS   = "HIVE_PAYMENT"

	HIVE_STORAGE = "HIVE_STORAGE"
	HIVE_USERS   = "HIVE_USERS"

	HIVE_TOKEN = "HIVE_TOKEN"

	HIVE_MEDIATION = "HIVE_MEDIATION"

	HIVE_JOBCREATOR = "HIVE_JOBCREATOR"

	HIVE_SOLVER = "HIVE_SOLVER"
)

const (
	NETWORK = "network"
)

const (
	WEB3_RPC_URL     = "WEB3_RPC_URL"
	WEB3_PRIVATE_KEY = "web3_private_key"

	HIVE_PRIVATE_KEY = "HIVE_PRIVATE_KEY"
	WEB3_CHAIN_ID    = "WEB3_CHAIN_ID"
)

const (
	RP_PRIVATE_KEY       = "RP_PRIVATE_KEY"
	JC_PRIVATE_KEY       = "JC_PRIVATE_KEY"
	SOLVER_PRIVATE_KEY   = "SOLVER_PRIVATE_KEY"
	MEDIATOR_PRIVATE_KEY = "MEDIATOR_PRIVATE_KEY"

	// 	deprecated cuz of options
)

/*Solver*/
const (
	SERVER_URL  = "server_url"
	SERVER_HOST = "server_host"
	SERVER_PORT = "server_port"
)

/*RP*/

const (
	OFFER_CPU   = "OFFER_CPU"
	OFFER_GPU   = "OFFER_GPU"
	OFFER_RAM   = "OFFER_RAM"
	OFFER_COUNT = "OFFER_COUNT"

	OFFER_MODULES = "OFFER_MODULES"
)
const (
	BACALHAU_API_HOST = "BACALHAU_API_HOST"
)

const PRICING_MODE = "PRICING_MODE"

/*JC*/

const (
	MODULE_NAME = "MODULE_NAME"
	MODULE_REPO = "MODULE_REPO"
	MODULE_HASH = "MODULE_HASH"
	MODULE_PATH = "MODULE_PATH"
)

/*Executor: Bacalhau*/

const (
	BACALHAU_BIN = "bacalhau_bin"
)
