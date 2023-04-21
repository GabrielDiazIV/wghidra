package rabbit

import (
	"context"
	"log"

	"github.com/gabrieldiaziv/wghidra/app/bo"
	"github.com/gabrieldiaziv/wghidra/app/system"
	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	REPLY_QUEUE    = "amq.rabbitmp.reply-to"
	REPLY_CONSUMER = "ReplyToConsumer"
)

func (r *rabbit) Close() {
	r.ch.Close()
	r.conn.Close()
}

func (r *rabbit) Consume(ctx context.Context, queue string, consumer string, autoAck bool) (<-chan amqp.Delivery, error) {
	return r.ch.Consume(
		queue,
		consumer,
		autoAck,
		false, //exclusive
		false, //no-local
		false, //no-wait
		nil,
	)
}

func (r *rabbit) ConsumeTask(ctx context.Context) <-chan amqp.Delivery {
	ch, err := r.Consume(
		ctx,
		r.chName,
		REPLY_CONSUMER,
		false, //autoAck
	)

	if err != nil {
		log.Fatalf("CONSUME TASK: %v", err)
	}

	return ch
}

func (r *rabbit) ConsumeResult(ctx context.Context) <-chan amqp.Delivery {
	ch, err := r.Consume(
		ctx,
		REPLY_QUEUE,
		REPLY_CONSUMER,
		true, //autoAck
	)

	if err != nil {
		log.Fatalf("CONSUME RESULT: %v", err)
	}

	return ch
}

func (r *rabbit) Publish(ctx context.Context, exchange string, routingKey string, pub amqp.Publishing) error {

	return r.ch.Publish(
		exchange,   // exchange
		routingKey, // channel
		false,      // mandatory
		false,      // immediate
		pub,        //value
	)
}

func (r *rabbit) ProduceTask(ctx context.Context, def bo.TaskDefinition) error {
	msg, err := system.Encode(def)
	if err != nil {
		return err
	}
	return r.Publish(
		ctx,
		r.exName, // exchange
		r.chName, // channel
		amqp.Publishing{
			ContentType: "application/json",
			Body:        msg,
			ReplyTo:     REPLY_QUEUE,
		},
	)
}

func (r *rabbit) ProduceResult(ctx context.Context, reciever string, res bo.TaskResult) error {
	msg, err := system.Encode(res)
	if err != nil {
		return err
	}
	return r.Publish(
		ctx,
		r.exName, // exchange
		reciever, // channel
		amqp.Publishing{
			ContentType: "application/json",
			Body:        msg,
		},
	)
}
