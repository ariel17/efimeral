package config

import "os"

const (
	ProductionEnv = "production"
	TestEnv       = "test"
)

var Environment string

func init() {
	Environment = os.Getenv("ENVIRONMENT")
}
