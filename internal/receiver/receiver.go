package receiver

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/periclescesar/event-processor/internal/application/services"
	amqp "github.com/rabbitmq/amqp091-go"
)

// EventConsumer consumes and processes events from a message broker.
type EventConsumer struct {
	eventService services.EventSaver // Service for saving events.
}

// NewEventConsumer creates a new instance of EventConsumer with the provided EventSaver.
func NewEventConsumer(eventService services.EventSaver) *EventConsumer {
	return &EventConsumer{eventService: eventService}
}

// Handle processes a message received from the message broker.
// It receives message and attempts to save it using the provided EventSaver.
// It returns an error if saving the event fails.
func (ec *EventConsumer) Handle(d amqp.Delivery) error {
	ctx := context.TODO()
	log.Debug("received a message...")
	log.Tracef("message to save: %s", d.Body)

	err := ec.eventService.Save(ctx, d.Body)
	if err != nil {
		return fmt.Errorf("saving event: %w", err)
	}

	return nil
}
