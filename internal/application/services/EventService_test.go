package services_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/periclescesar/event-processor/internal/application/services"

	"github.com/periclescesar/event-processor/internal/application/services/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestEventService_Save(t *testing.T) {
	ctx := context.TODO()
	validator := mocks.NewEventValidator(t)
	repo := mocks.NewEventRepository(t)

	es := services.NewEventService(validator, repo)

	validator.On("Validate", ctx, mock.Anything).Return(nil)
	repo.On("Save", ctx, mock.Anything).Return(nil)

	rawEvent := []byte(`{"eventType": "type", "tenantId": "123-123-123"}`)

	err := es.Save(ctx, rawEvent)
	assert.NoError(t, err)
}

func TestEventService_Save_UnmarshalErr(t *testing.T) {
	ctx := context.TODO()
	validator := mocks.NewEventValidator(t)
	repo := mocks.NewEventRepository(t)

	es := services.NewEventService(validator, repo)

	validator.AssertNotCalled(t, "Validate", ctx, mock.Anything)
	repo.AssertNotCalled(t, "Save", ctx, mock.Anything)

	rawEvent := []byte(`{"eventType": `)

	err := es.Save(ctx, rawEvent)
	assert.Error(t, err)
}

var errDummy = errors.New("dummy error")

func TestEventService_Save_InvalidEvent(t *testing.T) {
	ctx := context.TODO()
	validator := mocks.NewEventValidator(t)
	repo := mocks.NewEventRepository(t)

	es := services.NewEventService(validator, repo)

	validator.On("Validate", ctx, mock.Anything).Return(errDummy)
	repo.AssertNotCalled(t, "Save", ctx, mock.Anything)

	rawEvent := []byte(`{"eventType": "type"}`)

	err := es.Save(ctx, rawEvent)
	require.Error(t, err)
	assert.ErrorIs(t, err, errDummy)
}

func TestEventService_Save_ErrorOnSave(t *testing.T) {
	ctx := context.TODO()
	validator := mocks.NewEventValidator(t)
	repo := mocks.NewEventRepository(t)

	es := services.NewEventService(validator, repo)

	validator.On("Validate", ctx, mock.Anything).Return(nil)

	repo.On("Save", ctx, mock.Anything).Return(errDummy)

	rawEvent := []byte(`{"eventType": "type"}`)

	err := es.Save(ctx, rawEvent)
	require.Error(t, err)
	assert.ErrorIs(t, err, errDummy)
}
