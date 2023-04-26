package wghidra

import (
	"context"
	"io"

	"github.com/gabrieldiaziv/wghidra/app/bo"
)

func (w *wghidra) UploadRun(ctx context.Context, fstream io.ReadCloser, def bo.TaskDefinition) (bo.TaskResult, error) {

	panic("hehe")
}
