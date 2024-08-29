package main

import (
	"github.com/periclescesar/event-processor/configs"
	"github.com/streadway/amqp"
	"log"
)

func main() {
	configs.InitConfigs("deployments/.env")
	conn, err := amqp.Dial(configs.Rabbitmq.Uri)
	if err != nil {
		log.Fatalf("failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("failed to open a channel: %v", err)
	}
	defer ch.Close()

	body := `{ "eventType": "hello world" }`
	err = ch.Publish(
		"events.exchange",
		"",
		false,
		false,
		amqp.Publishing{
			Body: []byte(body),
		},
	)
	if err != nil {
		log.Fatalf("failed to publish a message: %v", err)
	}

	log.Printf("message sent: %v", body)
}
