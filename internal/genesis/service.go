package genesis

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Service struct {
	// TODO: add zerolog
	Log *logrus.Logger

	Conf *viper.Viper
}
