package sessions

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/ariel17/efimeral/api/apierrors"
	"github.com/ariel17/efimeral/api/config"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/mercadolibre/go-meli-toolkit/goutils/logger"
)

const (
	hubRepositoryName = "docker.io/ariel17/efimeral"
)

type Container struct {
	ID           string       `json:"id"`
	Distribution Distribution `json:"distribution"`
	Tag          string       `json:"tag"`
	CreatedAt    time.Time    `json:"created_at"`
}

type DockerClient interface {
	Pull(distribution Distribution, tag string) *apierrors.APIError
	Create(distribution Distribution, tag string) (*Container, *apierrors.APIError)
	Get(id string) (*Container, *apierrors.APIError)
	Destroy(id string) *apierrors.APIError
	List() ([]Container, *apierrors.APIError)
	Context() context.Context
}

type dockerClient struct {
	c       *client.Client
	context context.Context
}

func (dc *dockerClient) Pull(distribution Distribution, tag string) *apierrors.APIError {
	options := types.ImagePullOptions{}
	image := createImageName(distribution, tag)
	if _, err := dc.c.ImagePull(dc.Context(), image, options); err != nil {
		return apierrors.NewInternalServerError(err)
	}
	return nil
}

func (dc *dockerClient) Create(distribution Distribution, tag string) (*Container, *apierrors.APIError) {
	image := createImageName(distribution, tag)
	created, err := dc.c.ContainerCreate(dc.Context(), &container.Config{Image: image}, nil, nil, "")
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

	distribution, tag, apiErr := parseImageName(cJSON.Config.Image)
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
	}
	return &c, nil
}

func (dc *dockerClient) Destroy(id string) *apierrors.APIError {
	options := types.ContainerRemoveOptions{
		RemoveVolumes: true,
		RemoveLinks:   true,
		Force:         true,
	}
	if err := dc.c.ContainerRemove(context.Background(), id, options); err != nil {
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

func createImageName(distribution Distribution, tag string) string {
	return fmt.Sprintf("%s:%s-%s", hubRepositoryName, distribution, tag)
}

func parseImageName(image string) (Distribution, string, *apierrors.APIError) {
	aux := strings.Replace(image, hubRepositoryName, "", 1)
	aux = strings.Replace(aux, ":", "", 1)
	parts := strings.Split(aux, "-")
	if len(parts) != 2 {
		logger.Errorf("error parsing image name; has invalid parts: %s", nil, image)
		return "", "", apierrors.NewNotFoundError()
	}
	distribution := Distribution(parts[0])
	tags, found := availableDistributions[distribution]
	if !found {
		logger.Errorf("error parsing image name; distribution is invalid: %s", nil, image)
		return "", "", apierrors.NewNotFoundError()
	}
	found = false
	for _, tag := range tags {
		if tag == parts[1] {
			found = true
		}
	}
	if !found {
		logger.Errorf("error parsing image name; tag is not valid: %s", nil, image)
		return "", "", apierrors.NewNotFoundError()
	}

	return distribution, parts[1], nil
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
