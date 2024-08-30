package receiver

import (
	"encoding/json"
	"fmt"
	"github.com/periclescesar/event-processor/internal/event"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func EventConsumer(d amqp.Delivery) error {
	log.Printf("Received a message: %s", d.Body)

	ev := &event.Event{}
	err := json.Unmarshal(d.Body, ev)
	if err != nil {
		return fmt.Errorf("json decode: %w", err)
	}

	log.Printf("event type: %s", ev.EventType)
	// event decode
	// event saveUseCase
	return nil
}
