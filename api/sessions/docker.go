package sessions

import "github.com/ariel17/efimeral/api/apierrors"

type Container struct {
}

type DockerClient interface {
	Create(distribution Distribution, tag string) (*Container, apierrors.APIError)
	Destroy(container *Container) apierrors.APIError
	List() ([]*Container, apierrors.APIError)
}
