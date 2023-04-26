package cm

import (
	"github.com/docker/docker/client"
	"github.com/gabrieldiaziv/wghidra/app/bo/iface"
)

type containerManager struct {
	cli iface.DockerClient
}

func NewContainerManager(cli iface.DockerClient) iface.ContainerManager {
	return &containerManager{
		cli: cli,
	}
}

func NewDockerClient() (iface.DockerClient, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}

	return cli, nil
}
