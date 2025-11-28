package jsonschema

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/asiffer/puzzle/frontendtesting"
	"github.com/kaptinlin/jsonschema"
)

func TestGenerate(t *testing.T) {
	config, _ := frontendtesting.RandomNestedConfig()
	schema, err := Generate(config)
	if err != nil {
		t.Error(err)
	}
	bytes, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		t.Error(err)
	}

	compiler := jsonschema.NewCompiler()
	_, err = compiler.Compile(bytes)
	if err != nil {
		t.Error(fmt.Errorf("generated schema is invalid: %w", err))
	}
}
