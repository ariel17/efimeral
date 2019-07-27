package sessions

import (
	"time"

	"github.com/ariel17/efimeral/api/config"
	log "github.com/sirupsen/logrus"
)

// RunCleaner ticks for configured time steps for deleting older existing sessions.
func RunCleaner() {
	go cleaner()
}

func cleaner() {
	for range time.Tick(config.CleanTimeLapse) {
		containers, err := dc.List()
		if err != nil {
			log.Errorf("Cannot list existing containers: %v", err)
			if config.Environment != config.ProductionEnv {
				panic(err.Err)
			}
		}

		log.Infof("%d container(s) to check for expiration", len(containers))
		removed := 0
		for _, c := range containers {
			if c.HasExpired(config.MaxSessionDuration) {
				if err := dc.Destroy(c.ID); err != nil {
					log.Errorf("Cannot destroy container ID=%s: %v", c.ID, err)
					if config.Environment != config.ProductionEnv {
						panic(err.Err)
					}
				} else {
					log.Infof("Container ID=%s destroyed by expiration", c.ID)
					removed++
				}
			}
		}
		log.Infof("%d container(s) removed", removed)
	}
}
