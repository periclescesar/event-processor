package repository

import (
	"context"
	"fmt"
	"github.com/periclescesar/event-processor/internal/event"
	"go.mongodb.org/mongo-driver/mongo"
)

const collectionName = "events"

type MongoEventRepository struct {
	coll *mongo.Collection
}

func NewMongoEventRepository(db *mongo.Database) *MongoEventRepository {
	return &MongoEventRepository{coll: db.Collection(collectionName)}
}

func (e *MongoEventRepository) Save(ctx context.Context, ev *event.Event) error {
	_, err := e.coll.InsertOne(ctx, ev)
	if err != nil {
		return fmt.Errorf("save on mongo: %w", err)
	}
	return nil
}
