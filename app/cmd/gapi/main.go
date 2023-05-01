package main

import (
	"log"

	"github.com/gabrieldiaziv/wghidra/app/api/gapi"
	"github.com/gabrieldiaziv/wghidra/app/repo/cm"
	"github.com/gabrieldiaziv/wghidra/app/repo/dokr"
	"github.com/gabrieldiaziv/wghidra/app/repo/mock"
	"github.com/gabrieldiaziv/wghidra/app/srvc/wghidra"
)

// import "github.com/gabrieldiaziv/wghidra/app/api/gapi"

func main() {

	cli, err := cm.NewDockerClient()
	if err != nil {
		log.Fatalf("could not start client: %v", err)
	}

	api := gapi.NewGAPI(
		":6969",
		wghidra.NewWGhidra(
			dokr.NewRunner(cm.NewContainerManager(cli)),
			mock.NewStore(),
		))

	api.Start()
}
