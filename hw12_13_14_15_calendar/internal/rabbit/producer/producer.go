package producer

import (
	"context"
	"fmt"
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/logger"
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/rabbit"
	"github.com/rabbitmq/amqp091-go"
)

type Producer struct {
	*rabbit.RMQ
}

func NewProducer(uri, queue string) *Producer {
	return &Producer{rabbit.InitRabbitMQ(uri, queue)}
}

func (p *Producer) Publish(ctx context.Context, body []byte) (err error) {
	err = p.Chan.PublishWithContext(ctx,
		"",      // exchange
		p.Queue, // routing key
		false,   // mandatory
		false,   // immediate
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		return fmt.Errorf("failed to publish a message. Error: %s", err)
	}
	logger.Logger.Info(fmt.Sprintf("[x] Sent %s", body))
	return nil
}
