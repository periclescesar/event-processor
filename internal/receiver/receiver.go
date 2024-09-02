package receiver

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/periclescesar/event-processor/internal/event"
	"github.com/periclescesar/event-processor/internal/repository"
	schemaValidator "github.com/periclescesar/event-processor/pkg/schema-validator"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

type EventConsumer struct {
	validator *schemaValidator.SchemaValidator
	repo      *repository.MongoEventRepository
}

func NewEventConsumer(repo *repository.MongoEventRepository, validator *schemaValidator.SchemaValidator) *EventConsumer {
	return &EventConsumer{repo: repo, validator: validator}
}

func (ec *EventConsumer) Handle(d amqp.Delivery) error {
	ctx := context.TODO()
	log.Printf("Received a message: %s", d.Body)

	ev := &event.Event{}
	// event decode
	err := json.Unmarshal(d.Body, ev)
	if err != nil {
		return fmt.Errorf("event decode: %w", err)
	}

	// event validate

	errValid := ec.validator.Validate(ctx, ev.SchemaId(), d.Body)
	if errValid != nil {
		return errValid
	}

	// event save
	errSave := ec.repo.Save(ctx, ev)
	if errSave != nil {
		return errSave
	}
	return nil
}
