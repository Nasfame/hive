package genesis

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Service struct {
	// TODO: add zerolog
	Log logrus.Logger

	Conf *viper.Viper
}

func (s *Service) CreateContext() (ctx context.Context, cancel context.CancelFunc) {
	ctx, cancel = context.WithCancel(context.TODO())
	return
}
