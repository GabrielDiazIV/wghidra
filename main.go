package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/gabrieldiaziv/wghidra/app/bo"
	"github.com/gabrieldiaziv/wghidra/app/bo/iface"
	"github.com/gabrieldiaziv/wghidra/app/srvc/wghidra"
)

type mockdokr struct {
}

func NewMockDokr() iface.Dokr {
	return &mockdokr{}
}

func (d *mockdokr) Run(ctx context.Context, def bo.TaskDefinition) []bo.TaskResult {

	res := make([]bo.TaskResult, len(def.Tasks))
	for i, task := range def.Tasks {
		temp_res := bo.TaskResult{
			Name:  task.Name,
			Error: nil,
		}
		if task.Name == bo.DecompileTaskName {
			temp_res.Output = []bo.Function{
				{Name: "fn1", Parameters: []string{"p1", "p2"}, Body: "function body"},
				{Name: "fn2", Parameters: []string{"p1", "p2"}, Body: "function body222"},
			}
		} else if task.Name == bo.DissasmbleTaskName {
			temp_res.Output = "rax 1 2"
		} else if task.Name == bo.RunTaskName {
			temp_res.Output = "function output"
		} else {
			temp_res.Output = "ayyy yo"
		}

		res[i] = temp_res
	}

	return res
}

type mockstore struct {
	exe map[string][]byte
	dec map[string][]byte
}

func NewMockStore() iface.Store {
	return &mockstore{
		exe: make(map[string][]byte),
		dec: make(map[string][]byte),
	}
}

func (s *mockstore) PostExe(ctx context.Context, id string, stream io.Reader) (string, error) {

	data, err := io.ReadAll(stream)
	if err != nil {
		log.Fatal("could not read")
	}
	s.exe[id] = data
	return id, nil
}
func (s *mockstore) GetDecompiled(ctx context.Context, id string) (io.ReadCloser, error) {
	panic("not implemented") // TODO: Implement
}
func (s *mockstore) GetExe(ctx context.Context, id string) (io.ReadCloser, error) {
	data, ok := s.exe[id]
	if !ok {
		log.Fatal("could not find exe")
	}

	reader := io.NopCloser(bytes.NewReader(data))
	return reader, nil

}
func (s *mockstore) PostDecompiled(ctx context.Context, id string, stream io.Reader) (string, error) {
	panic("not implemented") // TODO: Implement
}

func main() {

	W := wghidra.NewWGhidra(
		NewMockDokr(),
		NewMockStore(),
	)

	ctx := context.Background()
	fakeReader := strings.NewReader("asdf")
	pid, fns, asm, err := W.ParseProject(ctx, fakeReader)
	if err != nil {
		fmt.Printf("parse err %v", err)
		return
	}

	fmt.Printf("project id: %s\n", pid)
	fmt.Printf("fns: %+v\n", fns)
	fmt.Printf("asm: %s\n", asm)

	// stPy := bo.ScriptTask{
	// 	ScriptName: "run.py",
	// 	Parameters: []string{"param.py", "p2", "p3"},
	// }
	// stJava := bo.ScriptTask{
	// 	ScriptName: "run.java",
	// 	Parameters: []string{"java", "j2", "j3"},
	// }
	// stRun := bo.ScriptTask{
	// 	ScriptName: bo.PyRunScript,
	// 	Parameters: []string{"prrr", "arg1", "arg2"},
	// }

}
