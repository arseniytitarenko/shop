package out

import amqp "github.com/rabbitmq/amqp091-go"

type BrokerSubscriber interface {
	Messages() <-chan amqp.Delivery
}
