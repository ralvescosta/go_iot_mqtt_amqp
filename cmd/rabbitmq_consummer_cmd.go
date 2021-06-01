package cmd

import (
	"iot/pkg/infra"
	"log"
	"os"
)

func StartQueueConsummer() {
	queue := infra.NewQueueConnection()

	err := queue.Connect()
	if err != nil {
		panic(err)
	}

	queueName := os.Getenv("DEFAULT_QUEUE")
	exchange := os.Getenv("DEFAULT_EXCHANGE")

	ch, err := queue.Consume(queueName, exchange)
	if err != nil {
		panic(err)
	}

	for d := range ch {
		log.Printf("AMQP SUB - Received Message: %s", d.Body)
	}
}
