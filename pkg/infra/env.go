package infra

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Panic("An error occurred reading the config file", err)
	}

	os.Setenv("MQTT_BROKER_URL", viper.GetString("MQTT_BROKER_URL"))
	os.Setenv("MQTT_BROKER_PORT", viper.GetString("MQTT_BROKER_PORT"))
	os.Setenv("MQTT_BROKER_PROTOCOL", viper.GetString("MQTT_BROKER_PROTOCOL"))
	os.Setenv("MQTT_BROKER_USER", viper.GetString("MQTT_BROKER_USER"))
	os.Setenv("MQTT_BROKER_PASSWORD", viper.GetString("MQTT_BROKER_PASSWORD"))
	os.Setenv("MQTT_DEFAULT_TOPIC", viper.GetString("MQTT_DEFAULT_TOPIC"))

	os.Setenv("AMQP_BROKER_URL", viper.GetString("AMQP_BROKER_URL"))
	os.Setenv("AMQP_BROKER_PORT", viper.GetString("AMQP_BROKER_PORT"))
	os.Setenv("AMQP_BROKER_USER", viper.GetString("AMQP_BROKER_USER"))
	os.Setenv("AMQP_BROKER_PASS", viper.GetString("AMQP_BROKER_PASS"))
	os.Setenv("DEFAULT_QUEUE", viper.GetString("DEFAULT_QUEUE"))
	os.Setenv("DEFAULT_EXCHANGE", viper.GetString("DEFAULT_EXCHANGE"))
}
