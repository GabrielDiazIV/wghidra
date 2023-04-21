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
	PostDecompiled(ctx context.Context, stream io.ReadCloser) (string, error)
}

type StoreProducer interface {
	PostExe(ctx context.Context, stream io.ReadCloser) (string, error)
	GetDecompiled(ctx context.Context, id string) (io.ReadCloser, error)
}
