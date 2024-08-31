package schema_validator

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/qri-io/jsonschema"
)

type SchemaValidator struct {
	rs *jsonschema.Schema
}

func NewSchemaValidator() *SchemaValidator {
	return &SchemaValidator{}
}

func (sv *SchemaValidator) ReadSchema() error {
	var schemaData = []byte(`{
  "$id": "file://configs/events-schemas/event-base.schema.json",
  "$schema": "https://json-schema.org/draft/2019-09/schema",
  "title": "base for all events",
  "description": "this schema should be used as a base for all events",
  "required": ["eventType"],
  "type": "object",
  "properties": {
    "eventType": {
      "type": "string",
      "minLength": 3,
      "maxLength": 255
    }
  }
}`)

	rs := &jsonschema.Schema{}
	if err := json.Unmarshal(schemaData, rs); err != nil {
		return fmt.Errorf("unmarshal schema: %w", err)
	}

	sv.rs = rs
	return nil
}

func (sv *SchemaValidator) Validate(ctx context.Context, json []byte) error {
	errs, err := sv.rs.ValidateBytes(ctx, json)
	if err != nil {
		return fmt.Errorf("validate: %w", err)
	}

	if len(errs) > 0 {
		return fmt.Errorf("validate: %w", errs[0])
	}

	return nil
}
