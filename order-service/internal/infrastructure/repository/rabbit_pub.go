package repository

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitPub struct {
	ch       *amqp.Channel
	exchange string
}

func NewRabbitPub(conn *amqp.Connection, exchange string) (*RabbitPub, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	err = ch.ExchangeDeclare(
		exchange,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &RabbitPub{ch: ch, exchange: exchange}, nil
}

func (r *RabbitPub) Publish(ctx context.Context, messageType string, payload string) error {
	return r.ch.PublishWithContext(ctx,
		r.exchange,
		messageType,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(payload),
			Type:        messageType,
		})
}
