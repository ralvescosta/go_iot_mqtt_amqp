package cmd

import (
	"iot/pkg/infra"
	"log"
	"os"
	"time"
)

func StartMQTTPublisher() {
	iot := infra.NewIoTClient()
	iot.Connect()
	topic := os.Getenv("MQTT_DEFAULT_TOPIC")

	for {
		log.Println("MQTT PUB - Pushing")
		iot.Pub(topic, 1, []byte("send to broker"))
		time.Sleep(time.Second * 2)
	}
}
