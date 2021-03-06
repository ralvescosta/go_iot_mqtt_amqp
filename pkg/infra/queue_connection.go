package infra

import (
	"fmt"
	"iot/pkg/app/interfaces"
	"os"

	"github.com/streadway/amqp"
)

type queueConnection struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func (q *queueConnection) Connect() error {
	broker := os.Getenv("AMQP_BROKER_URL")
	port := os.Getenv("AMQP_BROKER_PORT")
	user := os.Getenv("AMQP_BROKER_USER")
	password := os.Getenv("AMQP_BROKER_PASS")
	queue := os.Getenv("DEFAULT_QUEUE")

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s", user, password, broker, port))
	if err != nil {
		return err
	}

	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	_, err = ch.QueueDeclare(
		queue,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	q.ch = ch
	q.conn = conn

	return nil
}

func (q queueConnection) Produce(exchange, routingKey string, mandatory, immediate bool, data []byte) error {
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

func (q queueConnection) Consume(queue, routingKey string) (<-chan amqp.Delivery, error) {
	return q.ch.Consume(queue, routingKey, true, false, false, false, nil)
}

func (q queueConnection) Defer() {
	defer func() {
		q.ch.Close()
		q.conn.Close()
	}()
}

func NewQueueConnection() interfaces.IQueueConnection {
	return &queueConnection{}
}
