package services

import (
	"context"
	"github.com/periclescesar/event-processor/internal/application/services/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestEventService_Save(t *testing.T) {
	ctx := context.TODO()
	validator := mocks.NewEventValidator(t)
	repo := mocks.NewEventRepository(t)

	es := NewEventService(validator, repo)

	validator.On("Validate", ctx, mock.Anything).Return(nil)
	repo.On("Save", ctx, mock.Anything).Return(nil)

	rawEvent := []byte(`{"eventType": "type", "tenantId": "123-123-123"}`)

	err := es.Save(ctx, rawEvent)
	assert.NoError(t, err)
}
