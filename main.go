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
		go cmd.StartMQTTFakeIotDevice()
		go cmd.StartMQTTBridge()
		go cmd.StartDataConsummer()
	case "fake_iot_device":
		wg.Add(1)
		cmd.StartMQTTFakeIotDevice()
	case "mqtt_bridge":
		wg.Add(1)
		cmd.StartMQTTBridge()
	case "data_consumer":
		wg.Add(1)
		cmd.StartDataConsummer()
	default:
		panic(errors.New("you ned to config the operantion mode"))
	}

	wg.Wait()
}
