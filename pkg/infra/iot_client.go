package infra

import (
	"crypto/rand"
	"fmt"
	"iot/pkg/app/interfaces"
	"log"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type ioTCliente struct {
	client mqtt.Client
}

func (iot *ioTCliente) Connect() error {
	protocol := os.Getenv("MQTT_BROKER_PROTOCOL")
	broker := os.Getenv("MQTT_BROKER_URL")
	port := os.Getenv("MQTT_BROKER_PORT")

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("%s://%s:%s", protocol, broker, port))
	opts.SetClientID(guidV4())

	opts.OnConnect = iot.connectHandler
	opts.OnConnectionLost = iot.connectLostHandler

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	iot.client = client

	return nil
}

func guidV4() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

func (iot ioTCliente) Sub(topic string, qos int, Handler func(payload []byte, topic string)) {
	iot.client.Subscribe(topic, byte(qos), func(client mqtt.Client, msg mqtt.Message) {
		Handler(msg.Payload(), msg.Topic())
	})
}

func (iot ioTCliente) Pub(topic string, qos int, payload []byte) {
	iot.client.Publish(topic, byte(qos), false, payload)
}

func (iot ioTCliente) connectHandler(client mqtt.Client) {
	fmt.Println("Connected")
}

func (iot ioTCliente) connectLostHandler(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func NewIoTClient() interfaces.IIotCliente {
	return &ioTCliente{}
}
