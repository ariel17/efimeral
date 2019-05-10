package sessions

import (
	"encoding/json"
	"fmt"
	"time"
)

// Distribution refers to Linux distributions.
type Distribution string

const (
	ubuntu Distribution = "ubuntu"
	arch   Distribution = "arch"
)

var (
	availableDistributions = map[Distribution]struct{}{
		ubuntu: {},
		arch:   {},
	}
)

// SessionDistribution responds to the required data needed to create a new
// session.
type SessionDistribution struct {
	Distribution Distribution `json:"distribution"`
}

// UnmarshalJSON is a custom implementation for unmarshalling
// SessionDistribution in order to check valid distribution name.
func (sd *SessionDistribution) UnmarshalJSON(data []byte) error {
	type Alias SessionDistribution
	var aux Alias
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if _, found := availableDistributions[aux.Distribution]; found {
		sd.Distribution = aux.Distribution
	} else {
		return fmt.Errorf("'%s' is not available as distribution", aux.Distribution)
	}

	return nil
}

// Session represents a SO container instance on the Docker service.
type Session struct {
	ID string `json:"id"`
	SessionDistribution
	Tag       string     `json:"tag"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	Running   bool       `json:"running"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
