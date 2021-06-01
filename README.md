<div style="text-align:center">
  <img src="./docs/cover.png" />
</div>

# IoT With GoLang


## Get Started

- Requirements
  - GoLang 1.16.X
  - Docker
  - Docker Compose

- Configure the .env file in the root path

- Create a instance of MQTT Broker and RabbitMQ Broker

  ```bash
  docker-compose up
  ```

- Get Applications Packages

  ```
  go get -u
  ```

- Running application

  - Run all process
    ```bash
    make start
    ```

  - Ron one process per time
    ```bash
    make mqtt_pub
    ```

    ```bash
    make mqtt_sub
    ```

    ```bash
    make rabbitmq_consumer
    ```

## Description

In short, Internet of Things is interconnectivity between objects. To make this concept possible, have many technology behinds.

This project implements a simple cloud computer exemplification to works with IoT. Before presenting the adopted approach we need to understand what is MQTT? MQ Telemetry Transport - MQTT, is a lightweight messaging protocol for sensors and small mobile devices optimized for TCP / IP networks, if you want to improve on the subject read [more](https://developer.ibm.com/articles/iot-mqtt-why-good-for-iot/)

## Project Structure

<div align="center">
  <img src="./docs/fig.png" width="650"/>
</div>