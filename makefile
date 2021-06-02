start:
	GO_ENV=development OP_MODE=all go run main.go

fake_iot_device:
	GO_ENV=development OP_MODE=fake_iot_device go run main.go

mqtt_bridge:
	GO_ENV=development OP_MODE=mqtt_bridge go run main.go

data_consumer:
	GO_ENV=development OP_MODE=data_consumer go run main.go
