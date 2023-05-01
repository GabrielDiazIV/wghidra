package wghidra

import (
	"context"
	"errors"
	"io"

	"github.com/gabrieldiaziv/wghidra/app/bo"
	"github.com/gabrieldiaziv/wghidra/app/system"
	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
)

func (w *wghidra) PyRun(ctx context.Context, executeFunction string, paramters []string, functions []bo.Function) ([]bo.TaskResult, error) {

	funcs_json, err := system.Encode(functions_json{Functions: functions})
	if err != nil {
		log.Errorf("functions to json: %v", err)
		return nil, err
	}

	tarBuf, err := system.ToDockerTar(funcs_json, "input.json")
	if err != nil {
		return nil, err
	}

	// make argv {functionName, param1, param2, ...}
	argv := make([]string, len(paramters)+1)
	argv[0] = executeFunction
	for i, arg := range paramters {
		argv[i+1] = arg
	}

	id := uuid.New().String()
	def := bo.TaskDefinition{
		Tasks: []bo.UnitTask{bo.NewRunTask(id, argv)},
	}

	def.Exe = &tarBuf
	res := w.dokr.Run(ctx, def)
	return res, nil
}

// ParseProject implements defs.WGhidra
func (w *wghidra) ParseProject(ctx context.Context, fstream io.Reader) (string, []interface{}, error) {

	id := uuid.New().String()

	exeBuf, err := system.ToDockerTar(fstream, "input.out")
	if err != nil {
		return "", nil, err
	}

	readers := system.GetReaders(&exeBuf, 2)
	go func(rdr io.Reader) {
		_, err = w.store.PostExe(ctx, id, rdr)
		if err != nil {
			log.Errorf("could not upload exe: %v", err)
		}
	}(readers[0])

	defs := bo.TaskDefinition{
		Exe: readers[1],
		Tasks: []bo.UnitTask{
			bo.NewDecompileTask(id),
			// bo.NewDissasemblyTask(id),
		},
	}

	// run decompile task
	res := w.dokr.Run(ctx, defs)

	fns, okDec := res[0].Output["output"]
	if !okDec {
		log.Errorf("does not exist results: okASM = %b )", okDec)
		return "", nil, errors.New("not valid type")
	}

	fns_out, okDec := fns.([]interface{})
	if !okDec {
		log.Errorf("does not exist results: okASM = %b )", okDec)
		return "", nil, errors.New("not valid type")
	}

	return id, fns_out, nil
}

// RunScripts implements defs.WGhidra
func (w *wghidra) RunScripts(ctx context.Context, projectId string, def bo.TaskDefinition) ([]bo.TaskResult, error) {

	fstream, err := w.store.GetExe(ctx, projectId)
	if err != nil {
		log.Errorf("could not find project: %v", err)
		return nil, err
	}

	def.Exe = fstream
	defer fstream.Close()

	for i := range def.Tasks {
		def.Tasks[i].ID = bo.TaskID(projectId, def.Tasks[i].Name)
	}

	return w.dokr.Run(ctx, def), nil
}
