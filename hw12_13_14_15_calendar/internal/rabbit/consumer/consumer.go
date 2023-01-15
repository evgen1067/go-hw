package consumer

import (
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/rabbit"
	"github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	*rabbit.RMQ
}

func NewConsumer(uri, queue string) *Consumer {
	return &Consumer{rabbit.InitRabbitMQ(uri, queue)}
}

func (c *Consumer) Consume() (<-chan amqp091.Delivery, error) {
	return c.Chan.Consume(
		c.Queue, // queue
		"",      // consumer
		true,    // auto-ack
		false,   // exclusive
		false,   // no-local
		false,   // no-wait
		nil,     // args
	)
}
