package main

import (
	"github.com/CoopHive/hive/cmd"
	"log"
)

func main() {
	app := cmd.Hive()
	<-app.Done()

	log.Println("exiting gracefully")
}
