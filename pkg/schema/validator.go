package schema

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/periclescesar/event-processor/internal/application/event"
	"github.com/qri-io/jsonschema"
)

type Validator struct {
	registry    *jsonschema.SchemaRegistry
	schemasPath string
}

func NewSchemaValidator(registry *jsonschema.SchemaRegistry, schemasPath string) *Validator {
	return &Validator{registry: registry, schemasPath: schemasPath}
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

func (v *Validator) schemaIDBuilder(eventType string) string {
	pathFromRoot := strings.TrimLeft(v.schemasPath, "../")

	return fmt.Sprintf("file://%s/%s.schema.json", pathFromRoot, eventType)
}

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

	return nil
}
