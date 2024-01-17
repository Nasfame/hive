package config

import "os"

const COOPHIVE_MODULE_CONFIG_PATH = "/module.coophive"

var COOPHIVE_DATA_DIR = "/tmp/coophive/data"

func init() {
	coopHiveDataDir := os.Getenv("DATA_DIR")
	if coopHiveDataDir != "" {
		COOPHIVE_DATA_DIR = coopHiveDataDir
	}
}
