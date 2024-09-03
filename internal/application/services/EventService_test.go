package services_test

import (
	"context"
	"errors"
	"testing"

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

func TestEventService_Save_InvalidEvent(t *testing.T) {
	ctx := context.TODO()
	validator := mocks.NewEventValidator(t)
	repo := mocks.NewEventRepository(t)

	es := services.NewEventService(validator, repo)

	dummyError := errors.New("dummy error")

	validator.On("Validate", ctx, mock.Anything).Return(dummyError)
	repo.AssertNotCalled(t, "Save", ctx, mock.Anything)

	rawEvent := []byte(`{"eventType": "type"}`)

	err := es.Save(ctx, rawEvent)
	assert.Error(t, err)
	assert.ErrorAs(t, err, &dummyError)
}

func TestEventService_Save_ErrorOnSave(t *testing.T) {
	ctx := context.TODO()
	validator := mocks.NewEventValidator(t)
	repo := mocks.NewEventRepository(t)

	es := services.NewEventService(validator, repo)

	validator.On("Validate", ctx, mock.Anything).Return(nil)

	dummyError := errors.New("dummy error")
	repo.On("Save", ctx, mock.Anything).Return(dummyError)

	rawEvent := []byte(`{"eventType": "type"}`)

	err := es.Save(ctx, rawEvent)
	assert.Error(t, err)
	assert.ErrorAs(t, err, &dummyError)
}
