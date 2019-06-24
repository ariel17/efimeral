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
)

var (
	Environment string
	BaseURL     string
	Now         func() time.Time
	SentryDSN   string
)

func loadTestValues() {
	BaseURL = "http://localhost"
	Now = func() time.Time {
		return time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
	}
	SentryDSN = "https://15932fd4007f4e6d9e105442556df2b6@sentry.io/1488847"
}

func loadProductionValues() {
	BaseURL = os.Getenv("BASE_URL")
	Now = time.Now
	SentryDSN = "https://a07934d966db496eb3a27b003d6e9bfe@sentry.io/1488418"
}

func init() {
	if Environment = os.Getenv("ENVIRONMENT"); Environment == "" {
		Environment = "test"
	}
	log.Infof("Environment set to: %s", Environment)

	if Environment != ProductionEnv {
		loadTestValues()
	} else {
		loadProductionValues()
	}
}
