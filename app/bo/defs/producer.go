package defs

import (
	"context"

	"github.com/gabrieldiaziv/wghidra/app/bo"
)

type ProducerService interface {
	PublishTask(ctx context.Context, def bo.TaskDefinition) error
}
