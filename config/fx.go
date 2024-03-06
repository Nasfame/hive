package config

import (
	"log"
	"os"
	"slices"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"

	"github.com/CoopHive/hive/enums"
	"github.com/CoopHive/hive/utils"
)

var Module = fx.Options(
	fx.Provide(
		newConfig,
	),
	fx.Invoke(tempInitForFx),
	fx.WithLogger(func(conf *viper.Viper) (l fxevent.Logger) {
		if conf.GetBool(enums.DEBUG) {
			return &fxevent.ConsoleLogger{W: os.Stderr}
		}
		return fxevent.NopLogger
	}))

type out struct {
	fx.Out

	Conf *viper.Viper
}

func newConfig() (o out) {
	pf := pflag.NewFlagSet("conf", pflag.ContinueOnError)

	config := viper.New()

	checkDup := func(key string, block string) {
		if config.IsSet(key) {
			log.Fatalf("duplicate key found in config[%s]: %s", block, key)
		}
	}

	// formatEnvVar := func(key string) string {
	// 	k := strings.Replace("-", "_", key, -1)
	// 	k = strings.ToLower(k)
	// 	return k
	// }

	cmdFlags := map[string]bool{
		enums.APP_DIR:          true,
		enums.NETWORK:          true,
		enums.HIVE_CONTROLLER:  true,
		enums.DEALER:           true,
		enums.WEB3_PRIVATE_KEY: true,
		// enums.APP_DATA_DIR:   true,
		// enums.APP_PLUGIN_DIR: true,
		enums.BACALHAU_BIN: true,
	}

	for key, meta := range buildConfig {
		checkDup(key, "build")
		config.Set(key, meta.defaultVal)
	}

	for key, meta := range appConfig {
		checkDup(key, "app")

		config.SetDefault(key, meta.defaultVal)

		// automatic conversion of environment var key to `UPPER_CASE` will happen.
		config.BindEnv(key)

		if cmdFlags[key] {
			// key := strings.Replace("_", "-", key, -1)
			// read command-line arguments
			pf.String(key, meta.defaultVal, meta.desc)
			pflag.String(key, meta.defaultVal, meta.desc) // to show in usage
		}
	}

	osArgs := []string{}

	// bugfix: for `hive run cowsay:v0.1.2 -i Message="Hiro" --network sepolia` but defaulting to aurora
	// due to collusion with short hand vars
	for _, arg := range os.Args[1:] {
		// if strings.HasPrefix(arg, "--") {
		// 	osArgs = append(osArgs, arg)
		// }
		if !strings.HasPrefix(arg, "-") || strings.HasPrefix(arg, "--") {
			osArgs = append(osArgs, arg)
		}
	}

	if err := pf.Parse(osArgs); err != nil {
		logrus.Debugf("failed to parse args due to %v", err)
	}

	if err := config.BindPFlags(pf); err != nil {
		logrus.Debugf("failed to load flags:%v", err)
	}

	if config.GetBool(enums.DEBUG) {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.SetReportCaller(true)
	}

	appDir := config.GetString(enums.APP_DIR)
	logrus.Debugln("appDir: ", appDir)

	// appDataDir := config.GetString(enums.APP_DATA_DIR)
	// appDataDir = strings.Replace(appDataDir, AppDirSymbol, appDir, 1)
	// config.Set(enums.APP_DATA_DIR, appDataDir)

	/*Network related config*/

	network := config.GetString(enums.NETWORK)

	logrus.Debugln("network: ", network)

	if true {
		c, err := loadDApp(network)

		if err != nil {
			log.Fatal("failed to load the network related dApps")
		}

		for key, val := range c {
			key = strings.ToLower(key)
			curVal := config.GetString(key)

			defaultVal := ""

			if appConfig[key] != nil {
				defaultVal = appConfig[key].defaultVal
			}

			if curVal != "" && defaultVal != curVal {
				logrus.Debugf("key already set: %s", key)
				continue
			}
			logrus.Debugf("%v:%v\n", key, val)
			config.Set(key, val)
		}
		controller := config.Get(enums.HIVE_CONTROLLER)
		logrus.Debugln("controller: ", controller)
	}

	// override configurations with HIVE custom

	overRideMap := map[string]string{
		enums.HIVE_CHAIN_ID:    enums.WEB3_CHAIN_ID,
		enums.HIVE_RPC_URL:     enums.WEB3_RPC_URL,
		enums.HIVE_PRIVATE_KEY: enums.WEB3_PRIVATE_KEY,
		enums.HIVE_RPC_HTTP:    enums.WEB3_RPC_URL, // ordered: we don't use rpc http
		enums.HIVE_RPC_WS:      enums.WEB3_RPC_URL, // ordered
	}

	overrideKeysOrder := []string{
		enums.HIVE_CHAIN_ID,
		enums.HIVE_RPC_HTTP,
		enums.HIVE_RPC_URL,
		enums.HIVE_PRIVATE_KEY,
		enums.HIVE_RPC_WS,
	}
	overrideKeysOrder = slices.Compact(overrideKeysOrder) // remove dups

	if len(overRideMap) != len(overRideMap) {
		// not good enough: as there can be different entries in both or dup entries
		panic("entried not found not defined for some keys")
	}
	// check entries matches in both
	for _, key := range overrideKeysOrder {
		if _, ok := overRideMap[key]; !ok {
			panic("entry not found for key: " + key)
		}
	}

	// the below doesn't work as maps.Keys doesn't perseve order
	// if !slices.Equal(overrideKeysOrder, maps.Keys(overRideMap)) {
	// 	logrus.Errorf("%+v!=%+v", overRideMap, overrideKeysOrder)
	// 	panic("entries not matching for overrideKeys and order")
	// }

	for _, k := range overrideKeysOrder {
		v := overRideMap[k]
		overRiddingSetting := config.GetString(k)
		if overRiddingSetting != "" {

			smallK := strings.ToLower(k)

			// check websocket urls
			if strings.HasSuffix(smallK, "ws") {
				utils.PanicOnHTTPUrl(overRiddingSetting)
			}

			config.Set(v, overRiddingSetting)
			logrus.Debugf("overriding %s with %s", v, k)
		}
	}

	o.Conf = config
	return
}
