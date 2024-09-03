package schema

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

type Validator struct {
	registry *jsonschema.SchemaRegistry
}

func NewSchemaValidator() *Validator {
	return &Validator{registry: jsonschema.GetSchemaRegistry()}
}

func (v *Validator) loadSchemaFile(filePath string) (*jsonschema.Schema, error) {
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

func (v *Validator) RegistrySchemasFromPath(path string) error {
	files, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("retriving files on %s: %w", path, err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fullPath := filepath.Join(path, file.Name())
		schemaFile, errLoad := v.loadSchemaFile(fullPath)
		if errLoad != nil {
			return errLoad
		}

		schemaFile.Register("", v.registry)
	}

	return nil
}

func schemaIDBuilder(eventType string) string {
	return fmt.Sprintf("file://%s/%s.schema.json", SchemasPath, eventType)
}

func (v *Validator) Validate(ctx context.Context, event *event.Event) error {
	schema := v.registry.Get(ctx, schemaIDBuilder(event.EventType))
	if schema == nil {
		return fmt.Errorf("schema %s not found", event.EventType)
	}

	errs, err := schema.ValidateBytes(ctx, event.RawData)
	if err != nil {
		return fmt.Errorf("jsonschema validate: %w", err)
	}
	if len(errs) > 0 {
		return fmt.Errorf("jsonschema validate: %s", errs[0].Error())
	}

	return nil
}
