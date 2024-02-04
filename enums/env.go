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
	APP_DATA_DIR = "APP_DATA_DIR"
	MODULE_PATH  = "MODULE_PATH"
)
