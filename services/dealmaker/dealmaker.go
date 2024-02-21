package dealmaker

import (
	"context"
	"fmt"
	"path"
	"plugin"
	"runtime"
	"sync"

	"github.com/CoopHive/hive/enums"
	"github.com/CoopHive/hive/exp/dealer"
	"github.com/CoopHive/hive/internal/genesis"
)

type Service struct {
	name       string
	dealer     dealer.Dealer // can be a plugin
	ctx        context.Context
	cancelFunc context.CancelFunc

	m sync.Mutex

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
			func() {
				d.m.Lock()
				defer d.m.Unlock()
				f(dealID)
				d.Log.Debugf("Deal %s is agreed upon\n", dealID)
			}()

		case <-d.ctx.Done():

			d.Log.Printf("Context done. Exiting...")
			return
		}
	}
}

func (d *Service) setPlugin(plugin dealer.Dealer) {
	d.Log.Info("Setting plugin")
	d.dealer = plugin
}

func (d *Service) LoadPlugin(pluginName string) error {

	if !d.hasPluginSupport() {
		d.Log.Error("Plugins are not supported on this platform")
		return fmt.Errorf("loadplugin: is not supported on this platform")
	}

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

	customDealer := newFunction.(func(ctx context.Context) dealer.Dealer)(d.ctx)
	d.setPlugin(customDealer)

	return nil
}

func (d *Service) hasPluginSupport() bool {
	if runtime.GOOS == "windows" {
		return false
	}

	return true
}
