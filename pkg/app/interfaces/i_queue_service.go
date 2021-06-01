package interfaces

import "github.com/streadway/amqp"

type IQueueConnection interface {
	Connect() error
	Produce(exchange, routingKey string, mandatory, immediate bool, data []byte) error
	Consume(queue, routingKey string) (<-chan amqp.Delivery, error)
	Defer()
}
