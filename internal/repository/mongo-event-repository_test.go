package repository_test

import (
	"context"
	"testing"

	"github.com/periclescesar/event-processor/internal/repository"

	"github.com/periclescesar/event-processor/internal/application/event"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
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
		repo := repository.NewMongoEventRepository(dbMock)

		ev, _ := event.NewEventFromBytes([]byte(`{"eventType": "testType", "tenantId": "123a-asdf123-asdf123-asdf13"}`))

		err := repo.Save(context.TODO(), ev)
		assert.NoError(mt, err, "error not expected")
	})

	connMock.Run("error when saving an event", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})
		dbMock := mt.Client.Database(dbName)
		repo := repository.NewMongoEventRepository(dbMock)

		err := repo.Save(context.TODO(), &event.Event{EventType: "1234"})
		assert.Error(mt, err, "error expected")
	})
}
