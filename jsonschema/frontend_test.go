package jsonschema

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/asiffer/puzzle/frontendtesting"
	"github.com/asiffer/puzzle/jsonfile"
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

func testValidate(t *testing.T, i int) {
	// config
	config, _ := frontendtesting.RandomNestedConfig()
	// jsonschema
	schema, err := Generate(config)
	if err != nil {
		t.Error(err)
	}
	bytes, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		t.Error(err)
	}
	compiler := jsonschema.NewCompiler()
	compiler.SetAssertFormat(true)
	c, err := compiler.Compile(bytes)
	if err != nil {
		t.Error(fmt.Errorf("generated schema is invalid: %w", err))
	}

	// json data
	data, err := jsonfile.ToJSON(config)
	if err != nil {
		t.Error(err)
	}

	if result := c.ValidateJSON(data); !result.IsValid() {
		t.Log(string(bytes))
		t.Log(string(data))
		t.Error("data does not validate against schema", result.Errors)
	}
}

func FuzzTestValidate(f *testing.F) {
	for i := range 200 {
		f.Add(i)
	}
	f.Fuzz(testValidate)
}
