package jsonschema

import (
	"strings"

	"github.com/asiffer/puzzle"
)

const JsonSchemaFrontend puzzle.Frontend = "jsonschema"

func Generate(c *puzzle.Config) (*JSONSchema, error) {
	var root *JSONSchemaEntry
	schema := JSONSchema{
		Schema:          "https://json-schema.org/draft/2020-12/schema",
		JSONSchemaEntry: JSONSchemaEntry{Type: "object"},
	}

	for entry := range c.Entries() {
		key := entry.GetKey()
		root = &schema.JSONSchemaEntry
		keys := strings.Split(key, c.NestingSeparator)

		for len(keys) > 1 {
			if root.Properties == nil {
				root.Properties = make(map[string]*JSONSchemaEntry)
			}
			if props, exists := root.Properties[keys[0]]; exists {
				root = props
				root.Type = "object"
			} else {
				root.Properties[keys[0]] = &JSONSchemaEntry{
					Type: "object",
				}
				root = root.Properties[keys[0]]
			}
			keys = keys[1:]
		}

		s := JSONSchemaEntry{}
		if err := entry.Convert(JsonSchemaFrontend, &s); err != nil {
			return nil, err
		}
		if root.Properties == nil {
			root.Properties = make(map[string]*JSONSchemaEntry)
		}
		root.Properties[keys[0]] = &s
	}

	return &schema, nil
}
