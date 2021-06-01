start:
	GO_ENV=development GO_APP=all go run main.go

mqtt_pub:
	GO_ENV=development GO_APP=mqtt_pub go run main.go

mqtt_sub:
	GO_ENV=development GO_APP=mqtt_sub go run main.go

rabbitmq_consumer:
	GO_ENV=development GO_APP=rabbitmq_consumer go run main.go
