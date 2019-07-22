package sessions

import (
	"time"

	"github.com/ariel17/efimeral/api/config"
)

var (
	ch chan struct{}
)

func RunSessionCleaner() {
	go producer()
	go consumer()
}

func producer() {
	for range time.Tick(config.CleanTimeLapse) {
		ch <- struct{}{}
	}
}

func consumer() {
	for range ch {
		// TODO call Docker API to list active sessions
		// TODO remove older sessions
	}
}

func init() {
	ch = make(chan struct{})
}
