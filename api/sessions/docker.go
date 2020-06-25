package sessions

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/ariel17/efimeral/api/apierrors"
	"github.com/ariel17/efimeral/api/config"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	log "github.com/sirupsen/logrus"
)

const (
	hubRepositoryName = "docker.io/ariel17/efimeral"
	containerPort     = "4200"
)

type Container struct {
	ID           string              `json:"id"`
	Distribution config.Distribution `json:"distribution"`
	Tag          string              `json:"tag"`
	URL          string              `json:"url"`
	CreatedAt    time.Time           `json:"created_at"`
	DeletedAt    *time.Time          `json:"deleted_at,omitempty"`
}

func (c Container) HasExpired(d time.Duration) bool {
	running := config.Now().Sub(c.CreatedAt)
	return running > d
}

type DockerClient interface {
	Pull(distribution config.Distribution, tag string) *apierrors.APIError
	Create(distribution config.Distribution, tag string, cpus, memory int64, hostIP string, hostPort int) (*Container, *apierrors.APIError)
	Get(id string) (*Container, *apierrors.APIError)
	Destroy(id string) *apierrors.APIError
	List() ([]Container, *apierrors.APIError)
	Context() context.Context
}

type dockerClient struct {
	c       *client.Client
	context context.Context
}

func (dc *dockerClient) Pull(distribution config.Distribution, tag string) *apierrors.APIError {
	options := types.ImagePullOptions{}
	image := createImageName(distribution, tag)
	if _, err := dc.c.ImagePull(dc.Context(), image, options); err != nil {
		return apierrors.NewInternalServerError(err)
	}
	return nil
}

func (dc *dockerClient) Create(distribution config.Distribution, tag string, cpus, memory int64, hostIP string, port int) (*Container, *apierrors.APIError) {
	image := createImageName(distribution, tag)
	config := container.Config{Image: image}

	portBinding, apiErr := newPortBinding(hostIP, port)
	if apiErr != nil {
		return nil, apiErr
	}
	hostConfig := container.HostConfig{
		AutoRemove:   true,
		PortBindings: portBinding,
		Resources: container.Resources{
			Memory:     memory,
			MemorySwap: memory,
			CPUCount:   cpus,
		},
	}

	created, err := dc.c.ContainerCreate(dc.Context(), &config, &hostConfig, nil, "")
	if err != nil {
		return nil, apierrors.NewInternalServerError(err)
	}

	if err := dc.c.ContainerStart(dc.Context(), created.ID, types.ContainerStartOptions{}); err != nil {
		return nil, apierrors.NewInternalServerError(err)
	}

	return dc.Get(created.ID)
}

func (dc *dockerClient) Get(id string) (*Container, *apierrors.APIError) {
	cJSON, err := dc.c.ContainerInspect(dc.Context(), id)
	if err != nil {
		return nil, apierrors.NewInternalServerError(err)
	}

	if !cJSON.State.Running {
		return nil, nil
	}

	distribution, tag, apiErr := parseImageName(cJSON.Config.Image)
	if apiErr != nil {
		return nil, apiErr
	}

	hostPort, apiErr := extractHostPort(cJSON.NetworkSettings.Ports)
	if apiErr != nil {
		return nil, apiErr
	}

	createdAt, err := time.Parse(time.RFC3339, cJSON.Created)
	if err != nil {
		return nil, apierrors.NewInternalServerError(err)
	}

	c := Container{
		ID:           id,
		Distribution: distribution,
		Tag:          tag,
		CreatedAt:    createdAt,
		URL:          fmt.Sprintf("%s:%d", config.BaseURL, hostPort),
	}
	return &c, nil
}

func (dc *dockerClient) Destroy(id string) *apierrors.APIError {
	if err := dc.c.ContainerStop(dc.Context(), id, nil); err != nil {
		return apierrors.NewInternalServerError(err)
	}
	return nil
}

func (dc *dockerClient) List() ([]Container, *apierrors.APIError) {
	cl, err := dc.c.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		return nil, apierrors.NewInternalServerError(err)
	}

	containers := []Container{}
	for _, c := range cl {
		cc := Container{
			ID: c.ID,
		}
		containers = append(containers, cc)
	}

	return containers, nil
}

func (dc *dockerClient) Context() context.Context {
	return dc.context
}

func createImageName(distribution config.Distribution, tag string) string {
	return fmt.Sprintf("%s:%s-%s", hubRepositoryName, distribution, tag)
}

func parseImageName(image string) (config.Distribution, string, *apierrors.APIError) {
	aux := strings.Replace(image, hubRepositoryName, "", 1)
	aux = strings.Replace(aux, ":", "", 1)
	parts := strings.Split(aux, "-")
	if len(parts) != 2 {
		log.Errorf("error parsing image name; has invalid parts: %s", image)
		return "", "", apierrors.NewNotFoundError()
	}
	distribution := config.Distribution(parts[0])
	tags, found := config.AvailableDistributions[distribution]
	if !found {
		log.Errorf("error parsing image name; distribution is invalid: %s", image)
		return "", "", apierrors.NewNotFoundError()
	}
	found = false
	for _, tag := range tags {
		if tag == parts[1] {
			found = true
		}
	}
	if !found {
		log.Errorf("error parsing image name; tag is not valid: %s", image)
		return "", "", apierrors.NewNotFoundError()
	}

	return distribution, parts[1], nil
}

func extractHostPort(portMap nat.PortMap) (int, *apierrors.APIError) {
	var (
		hostPort int
		err      error
	)
	for k, v := range portMap {
		if k.Port() == containerPort {
			hostPort, err = strconv.Atoi(v[0].HostPort)
			if err != nil {
				return 0, apierrors.NewInternalServerError(err)
			}
			break
		}
	}
	if hostPort == 0 {
		err := errors.New("cannot find port mapping for container")
		return 0, apierrors.NewInternalServerError(err)
	}
	return hostPort, nil
}

func newPortBinding(hostIP string, hostPort int) (nat.PortMap, *apierrors.APIError) {
	hostBinding := nat.PortBinding{
		HostIP:   hostIP,
		HostPort: strconv.Itoa(hostPort),
	}

	containerPort, err := nat.NewPort("tcp", containerPort)
	if err != nil {
		return nil, apierrors.NewInternalServerError(err)
	}

	portBinding := nat.PortMap{
		containerPort: []nat.PortBinding{hostBinding},
	}

	return portBinding, nil
}

func newDockerClient() *dockerClient {
	c, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	return &dockerClient{
		c:       c,
		context: context.Background(),
	}
}

func NewDockerClient() DockerClient {
	if config.Environment == config.TestEnv {
		return newDockerMock()
	}
	return newDockerClient()
}
