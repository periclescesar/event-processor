package schema_validator

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/periclescesar/event-processor/internal/application/event"
	"github.com/qri-io/jsonschema"
)

const SchemasPath = "configs/events-schemas"

type SchemaValidator struct {
	registry *jsonschema.SchemaRegistry
}

func NewSchemaValidator() *SchemaValidator {
	return &SchemaValidator{registry: jsonschema.GetSchemaRegistry()}
}

func (sv *SchemaValidator) loadSchemaFile(filePath string) (*jsonschema.Schema, error) {
	byteFile, errR := os.ReadFile(filePath)
	if errR != nil {
		return nil, fmt.Errorf("read file %s: %w", filePath, errR)
	}

	rs := &jsonschema.Schema{}
	if errU := json.Unmarshal(byteFile, rs); errU != nil {
		return nil, fmt.Errorf("unmarshal schema: %w", errU)
	}

	return rs, nil
}

func (sv *SchemaValidator) RegistrySchemasFromPath(path string) error {
	files, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("retriving files on %s: %w", path, err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fullPath := filepath.Join(path, file.Name())
		schemaFile, errLoad := sv.loadSchemaFile(fullPath)
		if errLoad != nil {
			return errLoad
		}

		schemaFile.Register("", sv.registry)
	}

	return nil
}

func schemaIdBuilder(eventType string) string {
	return fmt.Sprintf("file://%s/%s.schema.json", SchemasPath, eventType)
}

func (sv *SchemaValidator) Validate(ctx context.Context, event *event.Event) error {
	schema := sv.registry.Get(ctx, schemaIdBuilder(event.EventType))
	if schema == nil {
		return fmt.Errorf("schema %s not found", event.EventType)
	}

	errs, err := schema.ValidateBytes(ctx, event.RawData)
	if err != nil {
		return fmt.Errorf("jsonschema validate: %s", err)
	}
	if len(errs) > 0 {
		return fmt.Errorf("jsonschema validate: %s", errs[0].Error())
	}

	return nil
}
