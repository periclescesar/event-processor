package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	conn *mongo.Client
	Db   *mongo.Database
}

var Manager = &Mongo{}

func Connect(ctx context.Context, uri, dbName string) error {
	conn, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return fmt.Errorf("mongo connection: %w", err)
	}

	Manager.conn = conn
	Manager.Db = conn.Database(dbName)
	return nil
}
