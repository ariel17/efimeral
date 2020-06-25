package sessions

import (
	"context"
	"fmt"

	"github.com/ariel17/efimeral/api/apierrors"
	"github.com/ariel17/efimeral/api/config"
)

type dockerMock struct {
	c   *Container
	err *apierrors.APIError
}

func (dc *dockerMock) Pull(distribution config.Distribution, tag string) *apierrors.APIError {
	return nil
}

func (dc *dockerMock) Create(distribution config.Distribution, tag string, cpus, memory int64, hostIP string, hostPort int) (*Container, *apierrors.APIError) {
	if dc.err != nil {
		return nil, dc.err
	}
	c := Container{
		ID:           "random-id",
		Distribution: distribution,
		Tag:          tag,
		CreatedAt:    config.Now().UTC(),
		URL:          fmt.Sprintf("%s:%d", config.BaseURL, 9000),
	}
	dc.c = &c
	return dc.c, nil
}

func (dc *dockerMock) Get(id string) (*Container, *apierrors.APIError) {
	if dc.err != nil {
		return nil, dc.err
	}
	return dc.c, nil
}

func (dc *dockerMock) Destroy(id string) *apierrors.APIError {
	return dc.err
}

func (dc *dockerMock) List() ([]Container, *apierrors.APIError) {
	return nil, nil
}

func (dc *dockerMock) Context() context.Context {
	return nil
}

func newDockerMock() *dockerMock {
	return &dockerMock{}
}
