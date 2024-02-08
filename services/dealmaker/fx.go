package dealmaker

import (
	"runtime"

	"github.com/spf13/viper"
	"go.uber.org/fx"

	"github.com/CoopHive/hive/enums"
	"github.com/CoopHive/hive/internal/genesis"
	"github.com/CoopHive/hive/services/dealmaker/internal"
)

var Module = fx.Options(
	fx.Provide(
		newServices,
	),
)

type in struct {
	fx.In
	*genesis.Service

	Conf *viper.Viper
}

type out struct {
	fx.Out
	DealerMaker *Service `name:"dealmaker"`
}

func newServices(i in) (o out) {
	conf := i.Conf

	dealerName := conf.GetString(enums.DEALER)
	// dealerPath := conf.GetString(enums.DEALER_PATH)

	s := newService(dealerName, i.Service)

	useDefaultPlugin := dealerName == "std-autoaccept" || runtime.GOOS == "windows"

CHOOSE_PLUGIN:
	if useDefaultPlugin {
		dealer := internal.NewAutoDealer(s.ctx)
		s.setPlugin(dealer)
		s.Log.Debugf("Using default plugin %s\n", dealerName)
	} else {
		err := s.loadPlugin(dealerName)
		if err != nil {
			s.Log.Errorf("Failed to load plugin %s: %v\n", dealerName, err)
			useDefaultPlugin = true
			goto CHOOSE_PLUGIN
		}
	}

	o = out{
		DealerMaker: s,
	}

	return
}
