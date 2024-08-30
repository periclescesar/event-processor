package main

import (
	"github.com/periclescesar/event-processor/configs"
	"github.com/periclescesar/event-processor/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func main() {
	configs.InitConfigs("deployments/.env")

	if err := rabbitmq.Connect(configs.Rabbitmq.Uri); err != nil {
		log.Fatalf("connection failure: %v", err)
	}

	errCons := rabbitmq.StartConsuming(func(d amqp.Delivery) {
		log.Printf("Received a message: %s", d.Body)
	})

	if errCons != nil {
		log.Fatalf("consuming failure: %v", errCons)
	}

	errC := rabbitmq.Close()
	if errC != nil {
		log.Fatalf("rabbitmq graceful shutdown: %v", errC)
	}
}
