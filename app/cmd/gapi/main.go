package main

import (
	"context"
	"log"
	"os"

	"github.com/gabrieldiaziv/wghidra/app/api/gapi"
	"github.com/gabrieldiaziv/wghidra/app/repo/cm"
	"github.com/gabrieldiaziv/wghidra/app/repo/dokr"
	"github.com/gabrieldiaziv/wghidra/app/repo/mock"
	"github.com/gabrieldiaziv/wghidra/app/srvc/wghidra"
	"github.com/gabrieldiaziv/wghidra/app/system"
)

// import "github.com/gabrieldiaziv/wghidra/app/api/gapi"

func main() {

	cli, err := cm.NewDockerClient()
	if err != nil {
		log.Fatalf("could not start client: %v", err)
	}

	file, err := os.Open("input.out")
	if err != nil {
		panic("could not load first script")
	}

	projectID := "my-project-id"
	buf, err := system.ToDockerTar(file, "input.out")
	if err != nil {
		panic("could not write to tar")
	}

	store := mock.NewStore()
	store.PostExe(context.Background(), projectID, &buf)

	api := gapi.NewGAPI(
		":6969",
		wghidra.NewWGhidra(
			dokr.NewRunner(cm.NewContainerManager(cli)),
			store,
		))

	api.Start()
}
