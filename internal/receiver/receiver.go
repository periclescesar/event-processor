package receiver

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/periclescesar/event-processor/configs"
	"github.com/periclescesar/event-processor/internal/event"
	"github.com/periclescesar/event-processor/internal/repository"
	"github.com/periclescesar/event-processor/pkg/mongodb"
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
	err = mongodb.Connect(context.TODO(), configs.Mongodb.Uri, "event-processor")
	repo := repository.NewMongoEventRepository(mongodb.Manager.Db)

	errSave := repo.Save(context.TODO(), ev)
	if errSave != nil {
		return errSave
	}
	return nil
}
