package main

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"

	"github.com/CoopHive/hive/cmd/hive"
)

func main() {
	hive.Execute()
}

func init() {
	//fmt.Printf("CoopHive: %s\n", hive.VERSION)

	err := godotenv.Load()
	if err != nil {
		log.Debug().Str("err", err.Error()).Msgf(".env not found")
	}

}
