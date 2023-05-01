package dokr

import (
	"github.com/gabrieldiaziv/wghidra/app/bo/iface"
)

type output_json struct {
	Output interface{} `json:"output,omitempty"`
}

type runner struct {
	containerManager iface.ContainerManager
}

func NewRunner(cm iface.ContainerManager) iface.Dokr {
	return &runner{
		containerManager: cm,
	}
}
