package infra

import (
	"os"
)

func init() {
	os.Setenv("MQTT_BROKER_URL", "test.mosquitto.org")
	os.Setenv("MQTT_BROKER_PORT", "1883")
	os.Setenv("MQTT_BROKER_PROTOCOL", "tcp")
	os.Setenv("MQTT_BROKER_USER", "")
	os.Setenv("MQTT_BROKER_PASSWORD", "")
	os.Setenv("MQTT_DEFAULT_TOPIC", "go_iot/mqtt/rabbitmq")

	os.Setenv("AMQP_BROKER_URL", "localhost")
	os.Setenv("AMQP_BROKER_PORT", "5672")
	os.Setenv("AMQP_BROKER_USER", "rabbitmq")
	os.Setenv("AMQP_BROKER_PASS", "123456")
	os.Setenv("DEFAULT_QUEUE", "temperature")
	os.Setenv("DEFAULT_EXCHANGE", "")

}
