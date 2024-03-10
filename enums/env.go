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

	GITHUB_REPO = "GITHUB_REPO"
)

const (
	APP_LOG_FILE_FORMAT = "APP_LOG_FILE_FORMAT"
	STD_REPO_URI        = "STD_REPO_URI"

	STD_MODULE_PREFIX = "STD_MODULE_PREFIX"

	STD_MODULE_FORMAT = STD_REPO_URI + STD_MODULE_PREFIX + "-%s"
)

const (
	APP_DIR = "app_dir"

	// Deprecated: and replaced by app_dir
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

	HIVE_MEDIATORS = "HIVE_MEDIATORS"

	HIVE_JOBCREATOR = "HIVE_JOBCREATOR"

	HIVE_SOLVER = "HIVE_SOLVER"
)

const (
	NETWORK = "network"
)

const (
	WEB3_RPC_URL = "WEB3_RPC_URL"
	HIVE_RPC_URL = "HIVE_RPC_URL" // TODO: add to web3 options, discuss with luke whether to backward support

	HIVE_RPC_WS   = "HIVE_RPC_WS"
	HIVE_RPC_HTTP = "HIVE_RPC_HTTP"

	WEB3_PRIVATE_KEY = "web3_private_key"

	HIVE_PRIVATE_KEY = "HIVE_PRIVATE_KEY"
	WEB3_CHAIN_ID    = "WEB3_CHAIN_ID"

	HIVE_CHAIN_ID = "HIVE_CHAIN_ID"
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

	// BACALHAU_REPO: set this when u r setting up multiple bacalhau clusters
	// in the same machine. For just a single machine just set the bacalhau repo.
	BACALHAU_REPO = "BACALHAU_REPO"
	BACALHAU_ENV  = "BACALHAU_ENV"

	// BACALHAU_ENVIRONMENT = "BACALHAU_ENVIRONMENT"

	// need to ensure this path is created
	BACALHAU_SERVE_IPFS_PATH = "BACALHAU_SERVE_IPFS_PATH"
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

	BACALHAU_RESULTS_DIR = "BACALHAU_RESULTS_DIR"

	BACALHAU_SPECS_DIR = "BACALHAU_SPECS_DIR"

	BACALHAU_JOBS_DIR = "BACALHAU_JOBS_DIR"

	REPO_DIR = "REPO_DIR"

	DOWNlOADS_DIR = "DOWNLOADS_DIR"
)

const (
	JOB_PRICE = "JOB_PRICE"
)
