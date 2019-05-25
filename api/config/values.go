package config

import (
	"os"
	"time"
)

const (
	ProductionEnv = "production"
	TestEnv       = "test"
)

var (
	Environment string
	Now         func() time.Time
)

func loadTestValues() {
	Now = func() time.Time {
		return time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
	}
}

func loadProductionValues() {
	Now = time.Now
}

func init() {
	Environment = os.Getenv("ENVIRONMENT")
	if Environment != ProductionEnv {
		loadTestValues()
	} else {
		loadProductionValues()
	}

}
