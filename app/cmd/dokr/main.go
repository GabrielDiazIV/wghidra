package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gabrieldiaziv/wghidra/app/bo"
	"github.com/gabrieldiaziv/wghidra/app/repo/dokr"
	yaml "gopkg.in/yaml.v2"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("need to pass file to parse")
		os.Exit(1)
	}

	def, err := readTaskDefinition(args[0])
	if err != nil {
		fmt.Println("invalid task file")
		os.Exit(1)
	}

	ctx := context.Background()
	runner, err := dokr.NewRunner(def)

	if err != nil {
		fmt.Printf("cannot create new runner: %v", err)
	}

	doneCh := make(chan bool)
	go runner.Run(ctx, doneCh)

	<-doneCh

}

func readTaskDefinition(fileName string) (bo.TaskDefinition, error) {

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return bo.TaskDefinition{}, err
	}

	var def bo.TaskDefinition
	err = yaml.Unmarshal(data, &def)
	if err != nil {
		return def, err
	}

	for i := 0; i < len(def.Tasks); i++ {
		reader, err := os.Open("tasks.tar.gz")
		if err != nil {
			panic("cannot find tar")
		}

		def.Tasks[i].Exe = reader
	}
	return def, nil
}
