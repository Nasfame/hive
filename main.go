package main

import (
	"github.com/CoopHive/hive/cmd/hive"
	"github.com/joho/godotenv"
)

func main() {
	hive.Execute()
}

func init() {
	//fmt.Printf("CoopHive: %s\n", hive.VERSION)

	err := godotenv.Load()
	if err != nil {
		//log.Debug().Str("err", err.Error()).Msgf(".env not found")
		//TODO: Doesn't look good, add custom flag
	}

}
