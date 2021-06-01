package cmd

import (
	"iot/pkg/handlers"
	"iot/pkg/infra"
	"os"
)

func StartMQTTSubscriber() {
	iot := infra.NewIoTClient()
	queue := infra.NewQueueConnection()

	err := iot.Connect()
	if err != nil {
		panic(err)
	}
	defer func() {
		iot.Defer()
	}()

	err = queue.Connect()
	if err != nil {
		panic(err)
	}
	defer func() {
		queue.Defer()
	}()

	topic := os.Getenv("MQTT_DEFAULT_TOPIC")
	iotHandler := handlers.IoTHandler{
		QueueConn: queue,
	}
	infinity := make(chan bool)

	iot.Sub(topic, 1, iotHandler.Handler)
	<-infinity
}
