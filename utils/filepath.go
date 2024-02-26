package utils

import (
	"os"
)

// EnsureDir: create the directory if doesn't exist
func EnsureDir(dir string) (p string, err error) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return dir, os.MkdirAll(dir, 0755)
	}
	return dir, nil
}
