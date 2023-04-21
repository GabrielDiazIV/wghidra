package rabbit

import (
	"fmt"
	"log"

	"github.com/gabrieldiaziv/wghidra/app/bo/iface"
	"github.com/gabrieldiaziv/wghidra/app/system"
	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	params = struct {
		name       string
		durable    bool
		autoDelete bool
		exclusive  bool
		noWait     bool
		args       amqp.Table
	}{
		name:       system.Unwrap(system.ENV.Rabbit.Chan),
		durable:    false,
		autoDelete: false,
		exclusive:  false,
		noWait:     false,
		args:       nil,
	}
)

type rabbit struct {
	ch     *amqp.Channel
	conn   *amqp.Connection
	chName string
	exName string
}

func NewRabbit() iface.Rabbit {
	ch, conn, err := declareQueue()
	if err != nil {
		log.Fatalf("RABBIT QUEUE: %v", err)
	}

	return &rabbit{
		ch:     ch,
		conn:   conn,
		chName: params.name,
		exName: system.Unwrap(system.ENV.Rabbit.Exch),
	}
}

func connString() string {
	return fmt.Sprintf("amqp:://%s:%s@%s:%s/",
		system.Unwrap(system.ENV.Rabbit.User),
		system.Unwrap(system.ENV.Rabbit.Pass),
		system.Unwrap(system.ENV.Rabbit.Host),
		system.Unwrap(system.ENV.Rabbit.Port),
	)
}

func declareQueue() (*amqp.Channel, *amqp.Connection, error) {
	conn, err := amqp.Dial(connString())
	if err != nil {
		log.Fatalf("RABBIT CONN: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil || ch != nil {
		log.Fatalf("RABBIT CHAN: %v", err)
	}

	_, err = ch.QueueDeclare(
		params.name,
		params.durable, params.autoDelete,
		params.exclusive,
		params.noWait,
		params.args,
	)

	return ch, conn, err
}
