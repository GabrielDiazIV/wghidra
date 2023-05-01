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

	def := bo.TaskDefinition{
		Tasks: []bo.UnitTask{bo.NewRunTask(argv)},
	}

	def.Exe = &tarBuf
	res := w.dokr.Run(ctx, def)
	return res, nil
}

// ParseProject implements defs.WGhidra
func (w *wghidra) ParseProject(ctx context.Context, fstream io.Reader) (string, []interface{}, string, error) {

	id := uuid.New().String()

	exeBuf, err := system.ToDockerTar(fstream, "input.out")
	if err != nil {
		return "", nil, "", err
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
			bo.NewDecompileTask(),
			bo.NewDissasemblyTask(),
		},
	}

	// run decompile task
	res := w.dokr.Run(ctx, defs)

	var idxDecompile int
	var idxAsm int

	// create decoder using result
	if res[0].Name == bo.DecompileTaskName {
		idxDecompile = 0
		idxAsm = 1
	} else {
		idxAsm = 0
		idxDecompile = 1
	}

	// TODO: UPDATE THIS ) VALUE
	asm, okAsm := res[idxAsm].Output["output"]
	fns, okDec := res[idxDecompile].Output["output"]

	if !okAsm || !okDec {
		log.Errorf("does not exist results: okASM = %b , okDec = %b)", okAsm, okDec)
		return "", nil, "", errors.New("not valid type")
	}

	// val, err := json.MarshalIndent(fns, "", "\t")
	// if err != nil {
	// 	log.Errorf("could not pretty json: %v", err)
	// 	return "", nil, "", errors.New("not valid type")
	// }

	asm_out, okAsm := asm.(string)
	fns_out, okDec := fns.([]interface{})

	if !okAsm || !okDec {
		log.Errorf("type of dec: %T", fns)
		log.Errorf("could not convert type: okASM = %b , okDec = %b)", okAsm, okDec)
		return "", nil, "", errors.New("not valid type")
	}

	return id, fns_out, asm_out, nil
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
