package constants

import "time"

var TopicTypeOut = "order.created"
var TopicTypeIn = "order.processed"
var ExchangeName = "events"
var QueueName = "order_status_processing"
var OutboxInterval = time.Second * 2
var ProcessingInterval = time.Second * 2
