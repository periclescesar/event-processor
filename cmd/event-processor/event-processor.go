package main

import (
	"context"
	"github.com/periclescesar/event-processor/configs"
	"github.com/periclescesar/event-processor/internal/receiver"
	"github.com/periclescesar/event-processor/pkg/mongodb"
	"github.com/periclescesar/event-processor/pkg/rabbitmq"
	"log"
)

func main() {
	configs.InitConfigs("deployments/.env")

	if err := rabbitmq.Connect(configs.Rabbitmq.Uri); err != nil {
		log.Fatalf("connection failure: %v", err)
	}

	err := mongodb.Connect(context.TODO(), configs.Mongodb.Uri, "event-processor")
	if err != nil {
		log.Fatalf("rabbitmq graceful shutdown: %v", err)
	}
	eventConsumer := receiver.NewEventConsumer(mongodb.Manager.Db)

	errCons := rabbitmq.StartConsuming(eventConsumer.Handle)

	if errCons != nil {
		log.Fatalf("consuming failure: %v", errCons)
	}

	errC := rabbitmq.Close()
	if errC != nil {
		log.Fatalf("rabbitmq graceful shutdown: %v", errC)
	}
}
