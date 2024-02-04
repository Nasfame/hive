package config

import (
	"github.com/joho/godotenv"
)

type configMap[T ~string] map[T]*argvMeta

type argvMeta struct {
	desc       string
	defaultVal string
}

func init() {
	// fmt.Printf("CoopHive: %s\n", hive.VERSION)

	err := godotenv.Load()
	if err != nil {
		// log.Debug().Str("err", err.Error()).Msgf(".env not found")
		// TODO: Doesn't look good, add custom flag
	}

}
