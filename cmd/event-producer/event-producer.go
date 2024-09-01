package main

import (
	"github.com/periclescesar/event-processor/configs"
	"github.com/periclescesar/event-processor/pkg/rabbitmq"
	"log"
	"os"
)

func main() {
	configs.InitConfigs("deployments/.env")

	if err := rabbitmq.Connect(configs.Rabbitmq.Uri); err != nil {
		log.Fatalf("connection failure: %v", err)
	}

	file, errR := os.ReadFile("test/mocked-events/user-created.json")
	if errR != nil {
		log.Fatalf("read event file: %v", errR)
	}
	body := string(file)

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
