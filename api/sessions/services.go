package sessions

import (
	"time"

	"github.com/ariel17/efimeral/api/apierrors"
)

var dc DockerClient

// NewSession builds a new container instance and returns its data associated
// or an error if it fails.
func NewSession(sr *SessionDistribution) (*Session, *apierrors.APIError) {
	container, err := dc.Create(sr.Distribution, sr.Tag)
	if err != nil {
		return nil, err
	}
	s := Session{
		Container: *container,
	}
	return &s, nil
}

// FetchSession returns the session data for given session ID, or an error if
// it fails to fetch.
func FetchSession(id string) (*Session, *apierrors.APIError) {
	container, err := dc.Get(id)
	if err != nil {
		return nil, err
	}
	if container == nil {
		return nil, nil
	}
	s := Session{
		Container: *container,
	}
	return &s, nil
}

// RemoveSession terminates and removes all data about given session ID. It
// returns its data at the delete moment or an error if the operation fails.
func RemoveSession(id string) (*Session, *apierrors.APIError) {
	container, err := dc.Get(id)
	if err != nil {
		return nil, err
	}
	if container == nil {
		return nil, nil
	}

	if err := dc.Destroy(id); err != nil {
		return nil, err
	}
	now := time.Now().UTC()
	s := Session{
		Container: *container,
		DeletedAt: &now,
	}
	return &s, nil
}

func pullAvailableImages() *apierrors.APIError {
	for distribution, tags := range availableDistributions {
		for _, tag := range tags {
			if err := dc.Pull(distribution, tag); err != nil {
				return err
			}
		}
	}
	return nil
}

func init() {
	dc = NewDockerClient()
	if err := pullAvailableImages(); err != nil {
		panic(err)
	}
}
