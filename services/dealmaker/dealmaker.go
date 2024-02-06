package dealmaker

import (
	"context"
	"os"
	"os/signal"

	"github.com/CoopHive/hive/exp/dealer"
	"github.com/CoopHive/hive/internal/genesis"
)

type Service struct {
	name       string
	plugin     dealer.Dealer
	ctx        context.Context
	cancelFunc context.CancelFunc

	*genesis.Service
}

func (d *Service) Name() string {
	return d.name
}

func (d *Service) DealsMatched(dealID string) {
	d.plugin.DealsMatched(dealID)
}

func (d *Service) DealsAgreed(f func(dealID string)) {
RECV_AGREE_DEALS:
	for {
		select {

		case dealID, ok := <-d.plugin.DealsAgreed():

			if !ok {
				d.Log.Debug("Channel closed. Exiting...")
				break RECV_AGREE_DEALS
			}
			f(dealID)

			d.Log.Debug("Deal %s is agreed upon\n", dealID)
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
	d.plugin = plugin
}
