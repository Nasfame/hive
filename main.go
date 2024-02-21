package main

import (
	"github.com/joho/godotenv"

	"github.com/CoopHive/hive/cmd"
)

func main() {
	cmd.Hive()
}

func init() {
	// fmt.Printf("CoopHive: %s\n", hive.VERSION)

	err := godotenv.Load(".env")
	if err != nil {
		// This is the best place to load, rest of the places are path relative to .env
		// log.Debug().Str("err", err.Error()).Msgf(".env not found")
		// TODO: Doesn't look good, add custom flag
	}

}
