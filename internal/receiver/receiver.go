package receiver

import (
	"context"
	"fmt"
	"log"

	"github.com/periclescesar/event-processor/internal/application/services"
	amqp "github.com/rabbitmq/amqp091-go"
)

type EventConsumer struct {
	eventService *services.EventService
}

func NewEventConsumer(eventService *services.EventService) *EventConsumer {
	return &EventConsumer{eventService: eventService}
}

func (ec *EventConsumer) Handle(d amqp.Delivery) error {
	ctx := context.TODO()
	log.Printf("Received a message: %s", d.Body)

	err := ec.eventService.Save(ctx, d.Body)
	if err != nil {
		return fmt.Errorf("saving event: %w", err)
	}

	return nil
}
