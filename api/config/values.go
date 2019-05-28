package config

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	ProductionEnv = "production"
	TestEnv       = "test"
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
