package main

import (
	"github.com/periclescesar/event-processor/configs"
	"github.com/periclescesar/event-processor/pkg/rabbitmq"
	"log"
)

func main() {
	configs.InitConfigs("deployments/.env")

	if err := rabbitmq.Connect(configs.Rabbitmq.Uri); err != nil {
		log.Fatalf("connection failure: %v", err)
	}

	body := `{ "eventType": "hello world" }`

	err := rabbitmq.Publish("events.exchange", body)
	if err != nil {
		log.Fatalf("publish failure: %v", err)
	}

	log.Printf("message sent: %v", body)

	errC := rabbitmq.Close()
	if errC != nil {
		log.Fatalf("rabbitmq graceful shutdown: %v", errC)
	}
}
