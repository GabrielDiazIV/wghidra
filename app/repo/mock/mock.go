package mock

import (
	"bytes"
	"context"
	"errors"
	"io"
	"log"

	"github.com/gabrieldiaziv/wghidra/app/bo/iface"
)

type store struct {
	exe map[string][]byte
	dec map[string][]byte
}

func NewStore() iface.Store {
	return &store{
		exe: make(map[string][]byte),
		dec: make(map[string][]byte),
	}
}

func (s *store) PostExe(ctx context.Context, id string, stream io.Reader) (string, error) {

	data, err := io.ReadAll(stream)
	if err != nil {
		log.Fatal("could not read")
	}
	s.exe[id] = data
	return id, nil
}
func (s *store) GetDecompiled(ctx context.Context, id string) (io.ReadCloser, error) {
	panic("not implemented") // TODO: Implement
}
func (s *store) GetExe(ctx context.Context, id string) (io.ReadCloser, error) {
	data, ok := s.exe[id]
	if !ok {
		return nil, errors.New("could not exe")
	}

	reader := io.NopCloser(bytes.NewReader(data))
	return reader, nil

}
func (s *store) PostDecompiled(ctx context.Context, id string, stream io.Reader) (string, error) {
	panic("not implemented") // TODO: Implement
}
