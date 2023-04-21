package iface

import (
	"context"
	"io"

	"github.com/docker/docker/client"
	"github.com/gabrieldiaziv/wghidra/app/bo"
)

type ContainerManager interface {
	PullImage(ctx context.Context, image string) error
	CreateContainer(ctx context.Context, task bo.UnitTask) (string, error)
	StartContainer(ctx context.Context, id string) error
	CopyTarOutput(ctx context.Context, id string) (io.ReadCloser, error)
	WaitForContainer(ctx context.Context, id string) (bool, error)
	RemoveContainer(ctx context.Context, id string) error
}

type DockerClient interface {
	client.ImageAPIClient
	client.ContainerAPIClient
}
