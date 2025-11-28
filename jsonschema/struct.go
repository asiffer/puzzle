package jsonschema

import "encoding/json"

type JSONSchema struct {
	Schema string `json:"$schema"`
	Id     string `json:"$id,omitempty"`
	JSONSchemaEntry
}

type JSONSchemaEntry struct {
	Title                string                      `json:"title,omitempty"`
	Format               string                      `json:"format,omitempty"`
	Pattern              string                      `json:"pattern,omitempty"`
	ContentEncoding      string                      `json:"contentEncoding,omitempty"`
	Description          string                      `json:"description,omitempty"`
	Type                 string                      `json:"type,omitempty"`
	Properties           map[string]*JSONSchemaEntry `json:"properties,omitempty"`
	PatternProperties    map[string]*JSONSchemaEntry `json:"patternProperties,omitempty"`
	Required             []string                    `json:"required,omitempty"`
	Items                *JSONSchemaEntry            `json:"items,omitempty"`
	AdditionalProperties bool                        `json:"additionalProperties,omitempty"`
	OneOf                []*JSONSchemaEntry          `json:"oneOf,omitempty"`
	PropertyConstraint
}

type MinMax struct {
	Negative *uint64
	Positive *uint64
}

func (m *MinMax) MarshalJSON() ([]byte, error) {
	if m.Negative != nil {
		b, err := json.Marshal(*m.Negative)
		if err != nil {
			return nil, err
		}
		return append([]byte("-"), b...), nil
	} else if m.Positive != nil {
		return json.Marshal(*m.Positive)
	} else {
		return json.Marshal(nil)
	}
}

type PropertyConstraint struct {
	MinItems    *uint64 `json:"minItems,omitempty"`
	MaxItems    *uint64 `json:"maxItems,omitempty"`
	UniqueItems *bool   `json:"uniqueItems,omitempty"`
	Minimum     *MinMax `json:"minimum,omitempty"`
	Maximum     *MinMax `json:"maximum,omitempty"`
}
