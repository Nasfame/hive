package genesis

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.uber.org/fx"

	"github.com/CoopHive/hive/enums"
)

var Module = fx.Options(
	fx.Provide(
		New,
	),
)

type in struct {
	fx.In
	Conf *viper.Viper
}

type out struct {
	fx.Out
	*Service
}

func New(i in) (o out) {

	conf := i.Conf

	logger := &logrus.Logger{
		Out:          os.Stderr,
		Formatter:    new(logrus.TextFormatter),
		Hooks:        make(logrus.LevelHooks),
		Level:        logrus.InfoLevel,
		ReportCaller: true,
	}

	if conf.GetBool(enums.DEBUG) {
		// logger.SetLevel(logrus.DebugLevel)
		logger.SetLevel(logrus.TraceLevel)
	}

	o = out{
		Service: &Service{
			logger,
			i.Conf,
		},
	}

	return
}
