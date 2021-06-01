package handlers

import (
	"iot/pkg/app/interfaces"
	"log"
	"os"
)

type IoTHandler struct {
	QueueConn interfaces.IQueueConnection
}

func (iot *IoTHandler) Handler(payload []byte, topic string) {
	log.Printf("MQTT SUB - topic: %s, msg: %s", topic, payload)

	queue := os.Getenv("DEFAULT_QUEUE")
	exchange := os.Getenv("DEFAULT_EXCHANGE")

	iot.QueueConn.Produce(exchange, queue, false, false, payload)
}
