package producer

import (
	"github.com/gabrieldiaziv/wghidra/app/bo"
	"github.com/gabrieldiaziv/wghidra/app/bo/defs"
	"github.com/gabrieldiaziv/wghidra/app/bo/iface"
)

type producerSrvc struct {
	pstore  iface.StoreProducer
	prabbit iface.RabbitProducer
}

func NewProducerSrvc(pstore iface.StoreProducer, prabbit iface.RabbitProducer) defs.ProducerService {
	srvc := &producerSrvc{
		pstore:  pstore,
		prabbit: prabbit,
	}

	return srvc
}
