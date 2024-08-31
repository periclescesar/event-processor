package receiver

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/periclescesar/event-processor/internal/event"
	"github.com/periclescesar/event-processor/internal/repository"
	schemaValidator "github.com/periclescesar/event-processor/pkg/mongodb/schema-validator"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type EventConsumer struct {
	repo *repository.MongoEventRepository
}

func NewEventConsumer(mongoDb *mongo.Database) *EventConsumer {
	repo := repository.NewMongoEventRepository(mongoDb)

	return &EventConsumer{repo: repo}
}

func (ec *EventConsumer) Handle(d amqp.Delivery) error {
	ctx := context.TODO()
	log.Printf("Received a message: %s", d.Body)

	ev := &event.Event{}
	// event decode
	err := json.Unmarshal(d.Body, ev)
	if err != nil {
		return fmt.Errorf("json decode: %w", err)
	}

	// event validate
	sv := schemaValidator.NewSchemaValidator()
	errSchema := sv.ReadSchema()
	if errSchema != nil {
		return errSchema
	}

	errValid := sv.Validate(ctx, d.Body)
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
