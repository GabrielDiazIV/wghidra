package defs

import (
	"context"
	"io"

	"github.com/gabrieldiaziv/wghidra/app/bo"
)

type WGhidra interface {
	UploadProject(ctx context.Context, fstream io.ReadCloser) (string, error)
	RunScripts(ctx context.Context, projectId string, def bo.TaskDefinition) (bo.TaskResult, error)
	UploadRun(ctx context.Context, fstream io.ReadCloser, def bo.TaskDefinition) (bo.TaskResult, error)
}
