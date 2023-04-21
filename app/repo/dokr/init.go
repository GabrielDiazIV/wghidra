package dokr

import (
	"github.com/docker/docker/client"
	"github.com/gabrieldiaziv/wghidra/app/bo"
	"github.com/gabrieldiaziv/wghidra/app/bo/iface"
	"github.com/gabrieldiaziv/wghidra/app/repo/cm"
)

type runner struct {
	def              bo.TaskDefinition
	containerManager iface.ContainerManager
}

func NewRunner(def bo.TaskDefinition) (iface.Dokr, error) {
	client, err := initDockerClient()
	if err != nil {
		return nil, err
	}

	return &runner{
		def:              def,
		containerManager: cm.NewContainerManager(client),
	}, nil
}

func initDockerClient() (iface.DockerClient, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}

	return cli, nil
}
