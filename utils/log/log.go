package __log

import (
	"os"

	_ "github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

var Log = &log.Logger{
	Out:       os.Stderr,
	Formatter: new(log.TextFormatter),
	Hooks:     make(log.LevelHooks),
	Level:     log.InfoLevel,
}

type Logger log.Logger

func New() (logger *Logger) {
	logger = (*Logger)(log.New())
	return
}
