package iface

import (
	"context"
	"io"
)

type Store interface {
	StoreProducer
	StoreWorker
}

type StoreWorker interface {
	GetExe(ctx context.Context, id string) (io.ReadCloser, error)
	PostDecompiled(ctx context.Context, id string, stream io.Reader) (string, error)
}

type StoreProducer interface {
	PostExe(ctx context.Context, id string, stream io.Reader) (string, error)
	GetDecompiled(ctx context.Context, id string) (io.ReadCloser, error)
}
