package dealmaker

import (
	"context"
	"fmt"
	"path"
	"plugin"
	"reflect"
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

	once bool

	dealsMatched map[string]bool // deals matched so far

	*genesis.Service
}

func (d *Service) Name() string {
	return d.name
}

func (d *Service) DealMatched(dealID string) {
	defer func() {
		if r := recover(); r != nil {
			d.Log.Errorf("Deal %s is matched but error occurred: %v", dealID, r)
			panic("dealer panic")
		}
	}()
	if _, matched := d.dealsMatched[dealID]; matched {
		return
	}
	d.dealer.DealMatched(dealID)
	d.dealsMatched[dealID] = true // TODO: is it the same dealID for different rps and same jc
}

// DealsAgreed should only be called exactly once
func (d *Service) DealsAgreed(f func(dealID string) error) {

	if d.once {
		panic("dealsAgreed should only be called once")
		return
	}

	defer func() {
		d.once = true
		if r := recover(); r != nil {
			d.Log.Fatalf("dealer paniced : %v", r)
		}
	}()

	doneDeals := map[string]bool{}

RECV_AGREE_DEALS:
	for {
		select {

		case dealID, ok := <-d.dealer.DealsAgreed():

			if !ok {
				d.Log.Debug("[dealer] Channel closed. Exiting...")
				break RECV_AGREE_DEALS
			}

			if doneDeals[dealID] {
				continue
			}

			func() {
				d.m.Lock()
				defer d.m.Unlock()

				defer func() {
					if r := recover(); r != nil {
						d.Log.Fatalf("f paniced : %v", r)
					}
				}()

				if err := f(dealID); err == nil {
					doneDeals[dealID] = true
					// d.dealsMatched[dealID] = true
					d.Log.Debugf("[dealer] Deal %s agreed tx", dealID)
				} else {
					d.Log.Errorf("[dealer] agreedDeal-%s agree failed due to %v", dealID, err)
				}
			}()

		case <-d.ctx.Done():

			d.Log.Printf("[dealer] Context done. Exiting...")
			return
		}
		d.Log.Debugf("total deals agreed so far: %d ; deals: %+v", len(doneDeals), reflect.ValueOf(doneDeals).MapKeys())

	}
}

// func (d *Service) Restart() {
// 	d.cancelFunc()
// 	d.ctx, d.cancelFunc = context.WithCancel(context.Background())
// 	d.m = sync.Mutex{}
// }

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
	d.Log.Infof("Loading plugin %s from %s", pluginName, pluginPath)
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
