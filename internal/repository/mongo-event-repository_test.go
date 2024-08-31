package repository

import (
	"context"
	"github.com/periclescesar/event-processor/internal/event"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"testing"
)

func TestEventRepository_Save(t *testing.T) {
	dbName := "event-processor-test"

	connMock := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	connMock.Run("success when saving an event", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "n", Value: 1},
			{Key: "acknowledgement", Value: true},
		})
		dbMock := mt.Client.Database(dbName)
		repo := NewMongoEventRepository(dbMock)

		err := repo.Save(context.TODO(), &event.Event{EventType: "testType"})
		assert.NoError(mt, err, "error not expected")
	})

	connMock.Run("error when saving an event", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})
		dbMock := mt.Client.Database(dbName)
		repo := NewMongoEventRepository(dbMock)

		err := repo.Save(context.TODO(), &event.Event{EventType: "1234"})
		assert.Error(mt, err, "error expected")
	})
}
