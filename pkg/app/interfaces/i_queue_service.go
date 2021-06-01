package interfaces

import "github.com/streadway/amqp"

type IQueueConnection interface {
	Connect() error
	Produce(exchange, routingKey string, mandatory, immediate bool, data []byte) error
	Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error)
	Defer()
}
