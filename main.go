package main

import (
	"errors"
	"iot/cmd"
	"os"
	"sync"
)

var wg sync.WaitGroup

func main() {

	mode := os.Getenv("OP_MODE")

	switch mode {
	case "all":
		wg.Add(3)
		go cmd.StartMQTTPublisher()
		go cmd.StartMQTTSubscriber()
		go cmd.StartQueueConsummer()
	case "mqtt_pub":
		cmd.StartMQTTPublisher()
	case "mqtt_sub":
		cmd.StartMQTTSubscriber()
	case "rabbitmq_consumer":
		cmd.StartQueueConsummer()
	default:
		panic(errors.New("you ned to config the operantion mode"))
	}

	wg.Wait()
}
