package schema_validator

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/qri-io/jsonschema"
	"os"
	"path/filepath"
)

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

func (sv *SchemaValidator) Validate(ctx context.Context, uri string, json []byte) error {
	schema := sv.registry.Get(ctx, uri)
	if schema == nil {
		return fmt.Errorf("schema %s not found", uri)
	}

	errs, err := schema.ValidateBytes(ctx, json)
	if err != nil {
		return fmt.Errorf("validate: %w", err)
	}

	if len(errs) > 0 {
		return fmt.Errorf("validate: %w", errs[0])
	}

	return nil
}
