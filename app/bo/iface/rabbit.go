package iface

import (
	"context"

	"github.com/gabrieldiaziv/wghidra/app/bo"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Rabbit interface {
	RabbitWorker
	RabbitProducer
	Close()
}

type RabbitWorker interface {
	ConsumeTask(ctx context.Context) <-chan amqp.Delivery
	ProduceResult(ctx context.Context, reciever string, res bo.TaskResult) error
	Close()
}
type RabbitProducer interface {
	ProduceTask(ctx context.Context, def bo.TaskDefinition) error
	ConsumeResult(ctx context.Context) <-chan amqp.Delivery
	Close()
}
