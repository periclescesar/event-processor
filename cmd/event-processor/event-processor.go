package main

import (
	"github.com/periclescesar/event-processor/configs"
	"github.com/periclescesar/event-processor/internal/receiver"
	"github.com/periclescesar/event-processor/pkg/rabbitmq"
	"log"
)

func main() {
	configs.InitConfigs("deployments/.env")

	if err := rabbitmq.Connect(configs.Rabbitmq.Uri); err != nil {
		log.Fatalf("connection failure: %v", err)
	}

	errCons := rabbitmq.StartConsuming(receiver.EventConsumer)

	if errCons != nil {
		log.Fatalf("consuming failure: %v", errCons)
	}

	errC := rabbitmq.Close()
	if errC != nil {
		log.Fatalf("rabbitmq graceful shutdown: %v", errC)
	}
}
