package defs

import (
	"context"
	"io"

	"github.com/gabrieldiaziv/wghidra/app/bo"
)

type WGhidra interface {
	ParseProject(ctx context.Context, fstream io.Reader) (string, []bo.Function, string, error)
	RunScripts(ctx context.Context, projectId string, def bo.TaskDefinition) ([]bo.TaskResult, error)
	PyRun(ctx context.Context, executeFunction string, functions []bo.Function) ([]bo.TaskResult, error)
}
