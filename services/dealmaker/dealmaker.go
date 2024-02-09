package dealmaker

import (
	"context"
	"os"
	"os/signal"
	"path"
	"plugin"

	"github.com/CoopHive/hive/enums"
	"github.com/CoopHive/hive/exp/dealer"
	"github.com/CoopHive/hive/internal/genesis"
)

type Service struct {
	name       string
	dealer     dealer.Dealer // can be a plugin
	ctx        context.Context
	cancelFunc context.CancelFunc

	*genesis.Service
}

func (d *Service) Name() string {
	return d.name
}

func (d *Service) DealMatched(dealID string) {
	defer func() {
		if r := recover(); r != nil {
			d.Log.Errorf("Deal %s is matched but error occurred: %v", dealID, r)
			panic("plugin paniced")
		}
	}()
	d.dealer.DealMatched(dealID)
}

func (d *Service) DealsAgreed(f func(dealID string)) {
	defer func() {
		if r := recover(); r != nil {
			d.Log.Fatalf("Critical error occurred: %v", r)
		}
	}()

RECV_AGREE_DEALS:
	for {
		select {

		case dealID, ok := <-d.dealer.DealsAgreed():

			if !ok {
				d.Log.Debug("Channel closed. Exiting...")
				break RECV_AGREE_DEALS
			}
			f(dealID)

			d.Log.Debugf("Deal %s is agreed upon\n", dealID)
		case <-d.ctx.Done():

			d.Log.Printf("Context done. Exiting...")
			return
		}
	}
}

func newService(name string, g *genesis.Service) *Service {
	ctx, cancelFunc := context.WithCancel(context.Background())

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	s := &Service{
		name,
		nil,
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

func (d *Service) setPlugin(plugin dealer.Dealer) {
	d.dealer = plugin
}

func (d *Service) loadPlugin(pluginName string) error {
	pluginPath := path.Join(d.Conf.GetString(enums.APP_PLUGIN_DIR), pluginName+".so")
	d.Log.Infof("Loading plugin %s from %s\n", pluginName, pluginPath)
	p, err := plugin.Open(pluginPath)
	if err != nil {
		return err
	}
	newFunction, err := p.Lookup("New")
	if err != nil {
		return err
	}

	d.dealer = newFunction.(func(ctx context.Context) dealer.Dealer)(d.ctx)

	return nil
}
