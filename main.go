package main

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"

	"github.com/bacalhau-project/lilypad/cmd/lilypad"
)

func main() {
	lilypad.Execute()
}

func init() {
	//fmt.Printf("Lilypad: %s\n", lilypad.VERSION)

	err := godotenv.Load()
	if err != nil {
		log.Debug().Str("err", err.Error()).Msgf(".env not found")
	}

}
