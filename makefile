start:
	GO_ENV=development OP_MODE=all go run main.go

mqtt_pub:
	GO_ENV=development OP_MODE=mqtt_pub go run main.go

mqtt_sub:
	GO_ENV=development OP_MODE=mqtt_sub go run main.go

rabbitmq_consumer:
	GO_ENV=development OP_MODE=rabbitmq_consumer go run main.go
