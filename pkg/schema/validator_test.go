package schema_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/qri-io/jsonschema"

	"github.com/periclescesar/event-processor/internal/application/event"

	"github.com/periclescesar/event-processor/pkg/schema"
	"github.com/stretchr/testify/assert"
)

func TestValidator_RegistrySchemasFromPath(t *testing.T) {

	t.Run("successfully registers schemas", func(t *testing.T) {
		v := schema.NewSchemaValidator(jsonschema.GetSchemaRegistry(), "../../configs/events-schemas")
		err := v.RegistrySchemas()
		assert.NoError(t, err)
	})

	t.Run("fails to read directory", func(t *testing.T) {
		v := schema.NewSchemaValidator(jsonschema.GetSchemaRegistry(), "invalid/path")
		err := v.RegistrySchemas()
		assert.Error(t, err)
	})
}

func TestValidator_Validate(t *testing.T) {
	v := schema.NewSchemaValidator(jsonschema.GetSchemaRegistry(), "../../configs/events-schemas")

	ctx := context.Background()

	errReg := v.RegistrySchemas()
	require.NoError(t, errReg)

	t.Run("successfully validate event", func(t *testing.T) {
		exampleEvent := &event.Event{
			EventType: "event-base",
			RawData:   []byte(`{"eventType": "event-base", "tenantId": "0dbf7fe2-8d30-4882-902e-bd0c48dc70ca"}`),
		}
		errValid := v.Validate(ctx, exampleEvent)
		require.NoError(t, errValid)
	})

	t.Run("fails wrong event type", func(t *testing.T) {
		exampleEvent := &event.Event{
			EventType: "example_event",
			RawData:   []byte(`{"key": "value"}`),
		}
		errValid := v.Validate(ctx, exampleEvent)
		require.Error(t, errValid)
	})

	t.Run("fails event invalid", func(t *testing.T) {
		exampleEvent := &event.Event{
			EventType: "event-base",
			RawData:   []byte(`{"key": "value"}`),
		}
		errValid := v.Validate(ctx, exampleEvent)
		assert.Error(t, errValid)
	})

	t.Run("fails event raw data is not a json", func(t *testing.T) {
		exampleEvent := &event.Event{
			EventType: "event-base",
			RawData:   []byte(`{"key":`),
		}
		errValid := v.Validate(ctx, exampleEvent)
		require.Error(t, errValid)
	})
}
