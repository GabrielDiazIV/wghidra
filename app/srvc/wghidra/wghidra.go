package wghidra

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"strings"

	"github.com/gabrieldiaziv/wghidra/app/bo"
	"github.com/gabrieldiaziv/wghidra/app/system"
	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
)

func (w *wghidra) PyRun(ctx context.Context, executeFunction string, paramters []string, functions string) (bo.TaskResult, error) {

	fnsReader := strings.NewReader(functions)
	tarBuf, err := system.ToTar(fnsReader, "functions.json")
	if err != nil {
		return bo.TaskResult{}, err
	}

	// make argv {functionName, param1, param2, ...}
	argv := make([]string, len(paramters)+1)
	argv[0] = executeFunction
	for i, arg := range paramters {
		argv[i+1] = arg
	}

	def := bo.TaskDefinition{
		Tasks: []bo.UnitTask{bo.NewRunTask("run", &tarBuf, argv)},
	}

	res := w.dokr.Run(ctx, def)
	return res[0], nil
}

// ParseProject implements defs.WGhidra
func (w *wghidra) ParseProject(ctx context.Context, fstream io.Reader) (string, []bo.Function, error) {

	// duplicate buffer
	var taskBuf bytes.Buffer
	storeReader := io.TeeReader(fstream, &taskBuf)

	// create id
	id := uuid.New().String()

	exeBuf, err := system.ToTar(storeReader, id)
	_, err = w.store.PostExe(ctx, id, &exeBuf)
	if err != nil {
		log.Errorf("could not upload exe: %v", err)
		return "", nil, err
	}

	// create decompile task
	defs := bo.TaskDefinition{
		Tasks: []bo.UnitTask{bo.NewDecompileTask("id", &taskBuf)},
	}

	// run decompile task
	res := w.dokr.Run(ctx, defs)

	// create decoder using result
	dec := json.NewDecoder(strings.NewReader(res[0].Output))

	// decode result
	var fns []bo.Function
	if err := dec.Decode(&fns); err != nil {
		log.Errorf("could not decode fns: %v", err)
		return "", nil, err
	}

	return id, fns, nil
}

// RunScripts implements defs.WGhidra
func (w *wghidra) RunScripts(ctx context.Context, projectId string, def bo.TaskDefinition) (bo.TaskResult, error) {

	fstream, err := w.store.GetExe(ctx, projectId)
	if err != nil {
		log.Errorf("could not find project: %v", err)
		return bo.TaskResult{}, err
	}

	defer fstream.Close()

}
