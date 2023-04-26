package main

import (
	"archive/tar"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/gabrieldiaziv/wghidra/app/bo"
	"github.com/gabrieldiaziv/wghidra/app/repo/cm"
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

	cli, err := cm.NewDockerClient()
	if err != nil {
		log.Fatalf("cannot make docker client %v", err)
	}

	runner := dokr.NewRunner(
		cm.NewContainerManager(cli),
	)

	results := runner.Run(ctx, def)
	for i, res := range results {
		if res.Error != nil {
			fmt.Printf("%d failed: %s\n", i, res.Error.Msg)
			continue
		}
		stream := tar.NewReader(res.TarStream)
		if _, err = stream.Next(); err != nil {
			fmt.Printf("%d failed: parse\n", i)
			continue
		}

		io.Copy(os.Stdout, stream)
	}

	if err != nil {
		panic(err)
	}

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
		reader, err := os.Open("./main.go")
		if err != nil {
			panic("cannot find tar")
		}

		def.Tasks[i].Exe = reader
	}
	return def, nil
}
