package main

import (
	"context"
	"fmt"
	"log"

	"github.com/periclescesar/event-processor/configs"
	"github.com/periclescesar/event-processor/internal/application/services"
	"github.com/periclescesar/event-processor/internal/receiver"
	"github.com/periclescesar/event-processor/internal/repository"
	"github.com/periclescesar/event-processor/pkg/mongodb"
	"github.com/periclescesar/event-processor/pkg/rabbitmq"
	schemaValidator "github.com/periclescesar/event-processor/pkg/schema-validator"
)

func main() {
	fmt.Println("starting event processor...")
	configs.InitConfigs()

	if err := rabbitmq.Connect(configs.Rabbitmq.URI); err != nil {
		log.Fatalf("connection failure: %v", err)
	}

	err := mongodb.Connect(context.TODO(), configs.Mongodb.URI, "event-processor")
	if err != nil {
		log.Fatalf("rabbitmq graceful shutdown: %v", err)
	}

	sv := schemaValidator.NewSchemaValidator()
	if sv.RegistrySchemasFromPath("configs/events-schemas") != nil {
		log.Fatalf("schema validation failure: %v", err)
	}

	repo := repository.NewMongoEventRepository(mongodb.Manager.Db)

	eventService := services.NewEventService(sv, repo)
	eventConsumer := receiver.NewEventConsumer(eventService)

	errCons := rabbitmq.StartConsuming(eventConsumer.Handle)

	if errCons != nil {
		log.Fatalf("consuming failure: %v", errCons)
	}

	errC := rabbitmq.Close()
	if errC != nil {
		log.Fatalf("rabbitmq graceful shutdown: %v", errC)
	}
}
