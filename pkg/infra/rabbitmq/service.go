package rabbitmq

import "github.com/streadway/amqp"

type QueueService struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func (q *QueueService) Connect() error {
	conn, err := amqp.Dial("amqp://rabbitmq:123456@localhost:5672")
	if err != nil {
		return err
	}

	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	_, err = ch.QueueDeclare(
		"temperature",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	return nil
}

func (q QueueService) Produce(exchange, routingKey string, mandatory, immediate bool, data []byte) error {
	err := q.ch.Publish(
		exchange,
		routingKey,
		mandatory,
		immediate,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        data,
		},
	)
	return err
}

func (q QueueService) Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error) {
	ch, err := q.ch.Consume(queue, consumer, autoAck, exclusive, noLocal, noWait, args)

	return ch, err
}

func (q QueueService) Defer() {
	defer func() {
		q.ch.Close()
		q.conn.Close()
	}()
}

func NewQueueService() *QueueService {
	return &QueueService{}
}
