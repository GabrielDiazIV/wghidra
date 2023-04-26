package dokr

import (
	"github.com/gabrieldiaziv/wghidra/app/bo/iface"
)

type runner struct {
	containerManager iface.ContainerManager
}

func NewRunner(cm iface.ContainerManager) iface.Dokr {
	return &runner{
		containerManager: cm,
	}
}
