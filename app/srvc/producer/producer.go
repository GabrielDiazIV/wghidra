package producer

import (
	"context"
	"net/http"

	"github.com/gabrieldiaziv/wghidra/app/bo"
)

func (p *producerSrvc) start() {
	for {
		select {}
	}

}

func (p *producerSrvc) PublishTask(ctx context.Context, def bo.TaskDefinition, r *http.Request) error {
	file, header, err := r.FormFile("file")
	if err != nil {
		return err
	}

}
