package config

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

// Distribution refers to Linux distributions.
type Distribution string

const (
	// ProductionEnv is the production environment name.
	ProductionEnv = "production"
	// TestEnv test environment name.
	TestEnv = "test"
	// DefaultCPUs is the default amount of CPUs to use in each container.
	DefaultCPUs = int64(1)
	// DefaultMemoryInBytes is the default amount of memory to use in each
	// container.
	DefaultMemoryInBytes              = int64(1024 * 1024 * 1024) // 1 GB
	ubuntu               Distribution = "ubuntu"
	arch                 Distribution = "archlinux"
)

var (
	// Environment is the current environment name.
	Environment string
	// BaseURL is the URL prefix where containers will be published.
	BaseURL string
	// Now is a time function holder.
	Now func() time.Time
	// SentryDSN is the value for Sentry to configure it (by environment).
	SentryDSN string
	// MaxSessionDuration is the maximum time allowed for each container.
	MaxSessionDuration time.Duration
	// CleanTimeLapse is the time between each execution for the cleaner job.
	CleanTimeLapse time.Duration
	// AvailableDistributions is the list of OS and tags availables in this
	// platform.
	AvailableDistributions map[Distribution][]string
)

func loadTestValues() {
	BaseURL = "http://localhost"
	Now = func() time.Time {
		return time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
	}
	SentryDSN = "https://15932fd4007f4e6d9e105442556df2b6@sentry.io/1488847"
	MaxSessionDuration = time.Minute
	CleanTimeLapse = 5 * time.Second
	AvailableDistributions = map[Distribution][]string{
		ubuntu: []string{"19.04"},
		arch:   []string{"latest"},
	}

}

func loadProductionValues() {
	BaseURL = os.Getenv("BASE_URL")
	Now = time.Now
	SentryDSN = "https://a07934d966db496eb3a27b003d6e9bfe@sentry.io/1488418"
	MaxSessionDuration = 20 * time.Minute
	CleanTimeLapse = 10 * time.Minute
	AvailableDistributions = map[Distribution][]string{
		ubuntu: []string{"20.04"},
		arch:   []string{"latest"},
	}

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

	log.Infof("Max session duration: %s", MaxSessionDuration.String())
	log.Infof("Cleaning time lapse: %s", CleanTimeLapse.String())
}
