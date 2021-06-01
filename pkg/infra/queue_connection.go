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
	broker := os.Getenv("MQTT_BROKER_URL")
	port := os.Getenv("MQTT_BROKER_PORT")
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

	return nil
}

func (q queueConnection) Produce(exchange, routingKey string, mandatory, immediate bool, data []byte) error {
	err := q.ch.Publish(
		exchange,
		routingKey,
		mandatory,
		immediate,
		amqp.Publishing{
			ContentType: "plain/text",
			Body:        data,
		},
	)
	return err
}

func (q queueConnection) Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error) {
	ch, err := q.ch.Consume(queue, consumer, autoAck, exclusive, noLocal, noWait, args)

	return ch, err
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
