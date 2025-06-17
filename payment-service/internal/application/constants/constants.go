package constants

import "time"

var TopicNameIn = "order.created"
var TopicNameOut = "order.processed"
var ExchangeName = "events"
var QueueName = "order_payment_processing"
var ProcessingInterval = time.Second * 2
var OutboxInterval = time.Second * 2
