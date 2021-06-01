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
	defer func() {
		iot.Defer()
	}()

	topic := os.Getenv("MQTT_DEFAULT_TOPIC")

	for {
		log.Println("MQTT PUB - Publishing")
		iot.Pub(topic, 1, []byte("send to broker"))
		time.Sleep(time.Second * 2)
	}

}
