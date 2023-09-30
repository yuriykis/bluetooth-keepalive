package device

import (
	"sync"

	log "github.com/sirupsen/logrus"
)

type LoggingMiddleware struct {
	next Devicer
}

func NewLoggingMiddleware(next Devicer) Devicer {
	return &LoggingMiddleware{
		next: next,
	}
}

func (lm *LoggingMiddleware) Up(wg *sync.WaitGroup) error {
	log.Infof("Device: %s", lm.next.String())
	defer func() {
		log.Infof("Device: %s is up", lm.next.String())
	}()
	return lm.next.Up(wg)
}

func (lm *LoggingMiddleware) String() string {
	return lm.next.String()
}
