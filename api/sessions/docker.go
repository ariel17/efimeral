package sessions

import (
	"context"
	"fmt"
	"time"

	"github.com/ariel17/efimeral/api/apierrors"
	"github.com/ariel17/efimeral/api/config"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
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
	return nil, nil
}

func (dc *dockerClient) Get(id string) (*Container, *apierrors.APIError) {
	return nil, nil
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
	return fmt.Sprintf("docker.io/ariel17/efimeral:%s-%s", distribution, tag)
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
