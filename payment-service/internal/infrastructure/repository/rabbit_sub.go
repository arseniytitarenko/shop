package repository

import amqp "github.com/rabbitmq/amqp091-go"

type RabbitSub struct {
	ch       *amqp.Channel
	messages <-chan amqp.Delivery
}

func NewRabbitSub(conn *amqp.Connection, exchange, queueName, bindingKey string) (*RabbitSub, error) {
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

	queue, err := ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	err = ch.QueueBind(
		queue.Name,
		bindingKey,
		exchange,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	messages, err := ch.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &RabbitSub{ch: ch, messages: messages}, nil
}

func (r *RabbitSub) Messages() <-chan amqp.Delivery {
	return r.messages
}
