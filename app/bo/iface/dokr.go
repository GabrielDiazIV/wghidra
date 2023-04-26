package iface

import (
	"context"

	"github.com/gabrieldiaziv/wghidra/app/bo"
)

type Dokr interface {
	Run(ctx context.Context, def bo.TaskDefinition) []bo.TaskResult
}
