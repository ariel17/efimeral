package sessions

import (
	"context"

	"github.com/ariel17/efimeral/api/apierrors"
)

type dockerMock struct {
}

func (dc *dockerMock) Pull(distribution Distribution, tag string) *apierrors.APIError {
	return nil
}

func (dc *dockerMock) Create(distribution Distribution, tag string) (*Container, *apierrors.APIError) {
	return nil, nil
}

func (dc *dockerMock) Destroy(container *Container) *apierrors.APIError {
	return nil
}

func (dc *dockerMock) List() ([]Container, *apierrors.APIError) {
	return nil, nil
}

func newDockerMock() *dockerMock {
	return &dockerMock{}
}

func (dc *dockerMock) Context() context.Context {
	return nil
}
