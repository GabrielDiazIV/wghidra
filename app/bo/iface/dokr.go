package iface

import "context"

type Dokr interface {
	Run(ctx context.Context, doneCh chan<- bool)
}
