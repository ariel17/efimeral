package config

// Distribution refers to Linux distributions.
type Distribution string

const (
	ubuntu Distribution = "ubuntu"
	arch   Distribution = "archlinux"
)

var (
	// AvailableDistributions contains the name and tags of distributions
	// offered for this platform.
	AvailableDistributions = map[Distribution][]string{
		ubuntu: []string{"19.04"},
		arch:   []string{"latest"},
	}
)
