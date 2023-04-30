package defs

import (
	"context"
	"io"

	"github.com/gabrieldiaziv/wghidra/app/bo"
)

type WGhidra interface {
	// UploadProject(ctx context.Context, fstream io.ReadCloser) (string, error)
	ParseProject(ctx context.Context, fstream io.Reader) (string, []bo.Function, error)
	RunScripts(ctx context.Context, projectId string, def bo.TaskDefinition) ([]bo.TaskResult, error)
	PyRun(ctx context.Context, executeFunction string, functions []bo.Function) (bo.TaskResult, error)
}
