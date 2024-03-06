package enums

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

type JobEntry struct {
	startTime time.Time
	endTime   time.Time
}

func (e *JobEntry) Duration() time.Duration {
	return e.endTime.Sub(e.startTime)
}

func (e *JobEntry) String() string {
	return fmt.Sprintf("start: %s, end: %s, duration: %s", e.startTime, e.endTime, e.Duration())
}

func (e *JobEntry) Start() {
	if !e.startTime.IsZero() {
		log.Debugf("start time set already")
		return
	}
	e.startTime = time.Now()
}

func (e *JobEntry) Stop() {
	if !e.endTime.IsZero() {
		log.Debugf("endTime set already")
		return
	}
	e.endTime = time.Now()
}
