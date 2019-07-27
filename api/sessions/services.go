package sessions

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/ariel17/efimeral/api/apierrors"
	"github.com/ariel17/efimeral/api/config"
	log "github.com/sirupsen/logrus"
)

const (
	creationRetries = 3 // due to already used ports
	minPort         = 1025
	maxPort         = 65535
	hostIP          = "0.0.0.0"
)

var (
	dc DockerClient
)

// NewSession builds a new container instance and returns its data associated
// or an error if it fails.
func NewSession(od *OSDistribution) (*Container, *apierrors.APIError) {
	var (
		c      *Container
		apiErr *apierrors.APIError
	)
	for retries := 0; retries < creationRetries; retries++ {
		rand.Seed(config.Now().UTC().UnixNano())
		hostPort := rand.Intn(maxPort-minPort) + minPort
		c, apiErr = dc.Create(od.Distribution, od.Tag, config.DefaultCPUs, config.DefaultMemoryInBytes, hostIP, hostPort)
		if apiErr == nil {
			return c, nil
		}
	}
	return nil, apiErr
}

// FetchSession returns the session data for given session ID, or an error if
// it fails to fetch.
func FetchSession(id string) (*Container, *apierrors.APIError) {
	return dc.Get(id)
}

// RemoveSession terminates and removes all data about given session ID. It
// returns its data at the delete moment or an error if the operation fails.
func RemoveSession(id string) (*Container, *apierrors.APIError) {
	container, err := dc.Get(id)
	if err != nil {
		return nil, err
	}
	if err := dc.Destroy(id); err != nil {
		return nil, err
	}
	now := config.Now().UTC()
	container.DeletedAt = &now
	return container, nil
}

func PullAvailableImages() *apierrors.APIError {
	log.Info("pre-fetching images ...")
	for distribution, tags := range availableDistributions {
		for _, tag := range tags {
			log.Infof("> %s:%s", distribution, tag)
			if err := dc.Pull(distribution, tag); err != nil {
				b, _ := json.Marshal(err)
				log.Error(fmt.Sprintf("error fetching image: %s", b))
				return err
			}
		}
	}
	log.Info("done :)")
	return nil
}

func init() {
	dc = NewDockerClient()
}
