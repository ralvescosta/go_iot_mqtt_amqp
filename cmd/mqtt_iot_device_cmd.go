package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"iot/pkg/app/dto"
	"iot/pkg/infra"
)

func StartMQTTFakeIotDevice() {
	iot := infra.NewIoTClient()
	iot.Connect()
	defer func() {
		iot.Defer()
	}()

	topic := os.Getenv("MQTT_DEFAULT_TOPIC")

	for {
		data := dto.IoTData{
			Device: fmt.Sprintf("%d", rand.Uint32()),
			Value:  rand.Intn(150),
			Time:   time.Now(),
		}
		toString, err := json.Marshal(data)
		if err != nil {
			log.Println("error when marshal struct")
			continue
		}

		log.Println("MQTT PUB - Publishing")
		iot.Pub(topic, 1, []byte(toString))
		time.Sleep(time.Second * 2)
	}

}
