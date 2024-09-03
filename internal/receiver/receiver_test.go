package receiver_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	mocks "github.com/periclescesar/event-processor/mocks/internal_/application/services"

	"github.com/periclescesar/event-processor/internal/receiver"

	amqp "github.com/rabbitmq/amqp091-go"
)

var errDummy = errors.New("dummy error")

func TestEventConsumer_Handle(t *testing.T) {
	delivery := amqp.Delivery{Body: []byte(`{"event": "test"}`)}

	t.Run("success", func(t *testing.T) {
		eventService := mocks.NewEventSaver(t)
		ec := receiver.NewEventConsumer(eventService)
		eventService.On("Save", mock.Anything, delivery.Body).Return(nil)

		err := ec.Handle(delivery)

		require.NoError(t, err)
	})
	t.Run("error", func(t *testing.T) {
		eventService := mocks.NewEventSaver(t)
		ec := receiver.NewEventConsumer(eventService)
		eventService.On("Save", mock.Anything, delivery.Body).Return(errDummy)

		err := ec.Handle(delivery)

		require.ErrorIs(t, err, errDummy)
	})
}
