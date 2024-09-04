// Package schema provides functionality for validating events against JSON schemas.
package schema

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/periclescesar/event-processor/internal/application/event"
	"github.com/qri-io/jsonschema"
)

// Validator validates events using JSON schemas.
type Validator struct {
	registry    *jsonschema.SchemaRegistry // Registry of JSON schemas.
	schemasPath string                     // Path to the directory containing schema files.
}

// NewSchemaValidator creates a new instance of Validator with the provided schema registry and path.
func NewSchemaValidator(registry *jsonschema.SchemaRegistry, schemasPath string) *Validator {
	return &Validator{registry: registry, schemasPath: schemasPath}
}

// loadSchemaFile reads a JSON schema file and unmarshals it into a jsonschema.Schema.
// It returns the schema and an error if the file could not be read or unmarshaled.
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

// RegistrySchemas loads and registers all JSON schemas from the specified directory.
// It reads schema files from the schemasPath and registers them in the schema registry.
func (v *Validator) RegistrySchemas() error {
	files, err := os.ReadDir(v.schemasPath)
	if err != nil {
		return fmt.Errorf("retrieving files on %s: %w", v.schemasPath, err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fullPath := filepath.Join(v.schemasPath, file.Name())
		schemaFile, errLoad := v.loadSchemaFile(fullPath)
		if errLoad != nil {
			return errLoad
		}

		schemaFile.Register("", v.registry)
	}

	return nil
}

// schemaIDBuilder constructs a schema ID based on the event type.
// It returns a schema ID in the format "file://path/to/schemas/eventType.schema.json".
func (v *Validator) schemaIDBuilder(eventType string) string {
	pathFromRoot := strings.TrimLeft(v.schemasPath, "../")

	return fmt.Sprintf("file://%s/%s.schema.json", pathFromRoot, eventType)
}

// Validate checks if the provided event data is valid according to its schema.
// It returns an error if the schema is not found or if validation fails.
func (v *Validator) Validate(ctx context.Context, event *event.Event) error {
	schema := v.registry.Get(ctx, v.schemaIDBuilder(event.EventType))
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

	log.Debug("jsonschema validated")

	return nil
}
