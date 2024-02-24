package config

import (
	"embed"
	"fmt"
	"path"
	"strings"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"github.com/CoopHive/hive/enums"
)

//go:embed dApps/*.env
var dApps embed.FS

const dAppFolderName = "dApps"

var NETWORKS = getNetworks()

func getNetworks() (networks []string) {
	dir, err := dApps.ReadDir(dAppFolderName)
	if err != nil {
		logrus.Fatalf("Error reading dApps directory: %v\n", err)
	}
	for _, file := range dir {
		networkFile := file.Name()

		if strings.HasSuffix(networkFile, ".env") {
			network := strings.TrimSuffix(networkFile, ".env")
			networks = append(networks, network)
		}
	}
	logrus.Debugln("support networks:", networks)
	return
}

func loadDApp(network string) (envMap map[string]string, err error) {

	dApp, _ := dApps.ReadFile(path.Join(dAppFolderName, fmt.Sprintf("%s.env", network)))

	envMap, err = godotenv.UnmarshalBytes(dApp)

	if err != nil {
		logrus.Errorf("Error loading %v.env file: due to %v \n", network, err)
		return
	}

	logrus.Debugln("envMap", envMap)

	isMediator := func(key string) bool {
		return strings.HasPrefix(key, strings.ToUpper(enums.HIVE_MEDIATORS))
	}

	var curMediators []string

	for key, value := range envMap {

		if isMediator(key) {
			logrus.Debugln("found mediator:", key, value)
			curMediators = append(curMediators, value)
			delete(envMap, key)
		}
	}

	if len(curMediators) == 0 {
		logrus.Fatalln("no mediators found")
	}

	envMap[enums.HIVE_MEDIATORS] = strings.Join(curMediators, ",")

	logrus.Debugln("mediation", envMap[enums.HIVE_MEDIATORS])

	return

}
