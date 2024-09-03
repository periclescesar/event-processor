package repository

import (
	"context"
	"fmt"

	"github.com/periclescesar/event-processor/internal/application/event"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

const collectionName = "events"

// MongoEventRepository implements EventRepository for MongoDB storage.
type MongoEventRepository struct {
	coll *mongo.Collection // MongoDB collection for storing events.
}

// NewMongoEventRepository creates a new MongoEventRepository instance.
func NewMongoEventRepository(db *mongo.Database) *MongoEventRepository {
	return &MongoEventRepository{coll: db.Collection(collectionName)}
}

// Save stores the given event in the MongoDB collection.
func (e *MongoEventRepository) Save(ctx context.Context, ev *event.Event) error {
	evPersist, errMap := ev.ToMap()
	if errMap != nil {
		return fmt.Errorf("convert event to map: %w", errMap)
	}

	id, err := e.coll.InsertOne(ctx, evPersist)
	if err != nil {
		return fmt.Errorf("save on mongo: %w", err)
	}

	log.Debugf("saved event with id: %s", id.InsertedID)

	return nil
}
