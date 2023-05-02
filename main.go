package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/gabrieldiaziv/wghidra/app/bo"
	"github.com/gabrieldiaziv/wghidra/app/repo/cm"
	"github.com/gabrieldiaziv/wghidra/app/repo/dokr"
	"github.com/gabrieldiaziv/wghidra/app/system"
)

func main() {
	cli, err := cm.NewDockerClient()
	if err != nil {
		panic(err)
	}

	runner := dokr.NewRunner(cm.NewContainerManager(cli))
	file, err := os.Open("input.out")
	if err != nil {
		panic(err)
	}

	tr, err := system.ToDockerTar(file, "input.out")
	if err != nil {
		panic(err)
	}

	def := bo.TaskDefinition{
		Exe:   &tr,
		Tasks: []bo.UnitTask{bo.NewDissasemblyTask("prjectid")},
	}

	res := runner.Run(context.Background(), def)

	b, err := json.MarshalIndent(res[0].Output, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Printf(string(b))

}
