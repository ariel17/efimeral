package config

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	ProductionEnv        = "production"
	TestEnv              = "test"
	DefaultCPUs          = int64(1)
	DefaultMemoryInBytes = int64(1024 * 1024 * 1024) // 1 GB
	SentryDSN            = "https://a07934d966db496eb3a27b003d6e9bfe:49b4d993cc74473bbf9d004edbef7e3c@sentry.io/1488418"
)

var (
	Environment string
	BaseURL     string
	Now         func() time.Time
)

func loadTestValues() {
	BaseURL = "http://localhost"
	Now = func() time.Time {
		return time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
	}
}

func loadProductionValues() {
	BaseURL = os.Getenv("BASE_URL")
	Now = time.Now
}

func init() {
	Environment = os.Getenv("ENVIRONMENT")
	log.Infof("ENVIRONMENT variable set to: %s", Environment)

	if Environment != ProductionEnv {
		loadTestValues()
	} else {
		loadProductionValues()
	}
}
