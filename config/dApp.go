package config

import (
	"embed"
	"fmt"
	"path"
	"strings"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"github.com/CoopHive/hive/enums"
	"github.com/CoopHive/hive/utils"
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

	networkFile := path.Join(dAppFolderName, fmt.Sprintf("%s.env", network))

	dApp, err := dApps.ReadFile(networkFile)

	if err != nil {
		panic(err)
	}

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

	switch network {

	case enums.AURORA:
		X_COOPHIVE_USER_HEADER = utils.Base64DecodeFast("WC1MaWx5cGFkLVVzZXI=")
		X_COOPHIVE_SIGNATURE_HEADER = utils.Base64DecodeFast("WC1MaWx5cGFkLVNpZ25hdHVyZQ==")

	case enums.HALCYON:
		X_COOPHIVE_USER_HEADER = "X-CoopHive-User"
		X_COOPHIVE_SIGNATURE_HEADER = "X-CoopHive-Signature"

	}

	return

}
