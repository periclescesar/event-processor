package repository

import (
	"context"
	"fmt"

	"github.com/periclescesar/event-processor/internal/application/event"
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
	_, err := e.coll.InsertOne(ctx, ev)
	if err != nil {
		return fmt.Errorf("save on mongo: %w", err)
	}

	return nil
}
