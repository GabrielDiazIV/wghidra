package main

import (
	"bytes"
	"context"
	"io"
	"log"

	"github.com/gabrieldiaziv/wghidra/app/bo/iface"
	"github.com/gabrieldiaziv/wghidra/app/repo/cm"
	"github.com/gabrieldiaziv/wghidra/app/repo/dokr"
	"github.com/gabrieldiaziv/wghidra/app/srvc/wghidra"
	"github.com/google/uuid"
)

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
	cli, err := cm.NewDockerClient()
	if err != nil {
		log.Fatalf("could not create cli: %v", err)
	}

	W := wghidra.NewWGhidra(
		dokr.NewRunner(cm.NewContainerManager(cli)),
		NewMockStore(),
	)

	ctx := context.Background()
	W.ParseProject(ctx)

}
