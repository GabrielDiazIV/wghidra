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

func (w *wghidra) PyRun(ctx context.Context, executeFunction string, paramters []string, functions string) ([]bo.TaskResult, error) {

	fnsReader := strings.NewReader(functions)
	tarBuf, err := system.ToTar(fnsReader, "functions.json")
	if err != nil {
		return nil, err
	}

	// make argv {functionName, param1, param2, ...}
	argv := make([]string, len(paramters)+1)
	argv[0] = executeFunction
	for i, arg := range paramters {
		argv[i+1] = arg
	}

	def := bo.TaskDefinition{
		Tasks: []bo.UnitTask{bo.NewRunTask(argv)},
	}

	def.Exe = &tarBuf
	res := w.dokr.Run(ctx, def)
	return res, nil
}

// ParseProject implements defs.WGhidra
func (w *wghidra) ParseProject(ctx context.Context, fstream io.Reader) (string, []bo.Function, string, error) {

	id := uuid.New().String()

	exeBuf, err := system.ToTar(fstream, id)
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
			bo.NewDecompileTask(),
			bo.NewDissasemblyTask(),
		},
	}

	// run decompile task
	res := w.dokr.Run(ctx, defs)

	var dec *json.Decoder
	var assembly string

	// create decoder using result

	if res[0].Name == bo.DecompileTaskName {
		dec = json.NewDecoder(strings.NewReader(res[0].Output))
		assembly = res[1].Output
	} else {
		dec = json.NewDecoder(strings.NewReader(res[1].Output))
		assembly = res[0].Output
	}

	// decode result
	var fns []bo.Function
	if err := dec.Decode(&fns); err != nil {
		log.Errorf("could not decode fns: %v", err)
		return "", nil, "", err
	}

	return id, fns, assembly, nil
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

	return w.dokr.Run(ctx, def), nil
}
