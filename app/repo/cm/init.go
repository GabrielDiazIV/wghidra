package cm

import "github.com/gabrieldiaziv/wghidra/app/bo/iface"

type containerManager struct {
	cli iface.DockerClient
}

func NewContainerManager(cli iface.DockerClient) iface.ContainerManager {
	return &containerManager{
		cli: cli,
	}
}
