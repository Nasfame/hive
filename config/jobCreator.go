package config

import "os"

const COOPHIVE_MODULE_CONFIG_PATH = "/lilypad_module.json.tmpl" //TODO: rebrand

var COOPHIVE_DATA_DIR = "/tmp/coophive/data"

func init() {
	coopHiveDataDir := os.Getenv("DATA_DIR")
	if coopHiveDataDir != "" {
		COOPHIVE_DATA_DIR = coopHiveDataDir
	}
}
