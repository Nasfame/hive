package dealmaker

import (
	"context"
	"os"
	"os/signal"
	"sync"

	"github.com/spf13/viper"
	"go.uber.org/fx"

	"github.com/CoopHive/hive/config"
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

	useDefaultPlugin := dealerName == config.DEFAULT_DEALER
	// I could have loaded the plugin indivitually in services but if every service will have a dealMakerService:
	// its easier to load here
	if !useDefaultPlugin {
		err := s.LoadPlugin(dealerName)
		if err != nil {
			s.Log.Errorf("Failed to load plugin %s: %v\n", dealerName, err)
		}
	}

	o = out{
		DealerMaker: s,
	}

	return
}

func newService(name string, g *genesis.Service) *Service {

	ctx, cancelFunc := signal.NotifyContext(context.Background(), os.Interrupt)

	s := &Service{
		name,
		internal.NewAutoDealer(ctx),
		ctx,
		cancelFunc,
		sync.Mutex{},
		false,
		g,
	}

	return s
}
