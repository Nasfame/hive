package dealmaker

import (
	"go.uber.org/fx"

	autoacceptdeals "github.com/CoopHive/hive/exp/dealer/autoAccept"
	"github.com/CoopHive/hive/internal/genesis"
)

var Module = fx.Options(
	fx.Provide(
		newServices,
	),
)

type in struct {
	fx.In
	*genesis.Service
}

type out struct {
	fx.Out
	DealerMaker *Service `name:"dealmaker"` // this should be fx.Annotated
}

func newServices(i in) (o out) {

	s := newService("autoaccept", i.Service)

	a := autoacceptdeals.New(s.ctx)

	s.setPlugin(a)

	o = out{
		DealerMaker: s,
	}

	return
}
