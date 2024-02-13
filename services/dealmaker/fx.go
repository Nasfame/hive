package dealmaker

import (
	"context"
	"os"
	"os/signal"

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

	/*	useDefaultPlugin := dealerName == "std-autoaccept" || runtime.GOOS == "windows"

		CHOOSE_PLUGIN:
			if !useDefaultPlugin {
				err := s.LoadPlugin(dealerName)
				if err != nil {
					s.Log.Errorf("Failed to load plugin %s: %v\n", dealerName, err)
					useDefaultPlugin = true
					goto CHOOSE_PLUGIN
				}
			}*/

	o = out{
		DealerMaker: s,
	}

	return
}

func newService(name string, g *genesis.Service) *Service {
	ctx, cancelFunc := context.WithCancel(context.Background())

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	s := &Service{
		name,
		internal.NewAutoDealer(ctx),
		ctx,
		cancelFunc,
		g,
	}

	go func() {
		sig := <-c
		s.Log.Errorf("Got signal:%s", sig) // TODO: use fx signals if possible
		cancelFunc()
	}()

	return s
}
