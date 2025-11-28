package jsonschema

import (
	"net"
	"time"

	"github.com/asiffer/puzzle"
)

func init() {
	puzzle.BoolConverter.Register(JsonSchemaFrontend, puzzle.ConvertCallbackFactory1(withDescription(boolJsonSchema)))
	puzzle.DurationConverter.Register(JsonSchemaFrontend, puzzle.ConvertCallbackFactory1(withDescription(durationJsonSchema)))
	puzzle.Float64Converter.Register(JsonSchemaFrontend, puzzle.ConvertCallbackFactory1(withDescription(float64JsonSchema)))
	puzzle.IntConverter.Register(JsonSchemaFrontend, puzzle.ConvertCallbackFactory1(withDescription(intJsonSchema)))
	puzzle.Int64Converter.Register(JsonSchemaFrontend, puzzle.ConvertCallbackFactory1(withDescription(int64JsonSchema)))
	puzzle.StringConverter.Register(JsonSchemaFrontend, puzzle.ConvertCallbackFactory1(withDescription(stringJsonSchema)))
	puzzle.UintConverter.Register(JsonSchemaFrontend, puzzle.ConvertCallbackFactory1(withDescription(uintJsonSchema)))
	puzzle.Uint64Converter.Register(JsonSchemaFrontend, puzzle.ConvertCallbackFactory1(withDescription(uint64JsonSchema)))
	puzzle.BytesConverter.Register(JsonSchemaFrontend, puzzle.ConvertCallbackFactory1(withDescription(bytesJsonSchema)))
	puzzle.IPConverter.Register(JsonSchemaFrontend, puzzle.ConvertCallbackFactory1(withDescription(ipJsonSchema)))
	puzzle.Int8Converter.Register(JsonSchemaFrontend, puzzle.ConvertCallbackFactory1(withDescription(int8JsonSchema)))
	puzzle.Int16Converter.Register(JsonSchemaFrontend, puzzle.ConvertCallbackFactory1(withDescription(int16JsonSchema)))
	puzzle.Int32Converter.Register(JsonSchemaFrontend, puzzle.ConvertCallbackFactory1(withDescription(int32JsonSchema)))
	puzzle.Uint8Converter.Register(JsonSchemaFrontend, puzzle.ConvertCallbackFactory1(withDescription(uint8JsonSchema)))
	puzzle.Uint16Converter.Register(JsonSchemaFrontend, puzzle.ConvertCallbackFactory1(withDescription(uint16JsonSchema)))
	puzzle.Uint32Converter.Register(JsonSchemaFrontend, puzzle.ConvertCallbackFactory1(withDescription(uint32JsonSchema)))
	puzzle.StringSliceConverter.Register(JsonSchemaFrontend, puzzle.ConvertCallbackFactory1(withDescription(stringSliceJsonSchema)))
	puzzle.Float32Converter.Register(JsonSchemaFrontend, puzzle.ConvertCallbackFactory1(withDescription(float32JsonSchema)))
}

const UZERO = uint(0)

// supported types

type jsonSchemaConverter[T any] = func(entry *puzzle.Entry[T], schema *JSONSchemaEntry) error

func withDescription[T any](fun jsonSchemaConverter[T]) jsonSchemaConverter[T] {
	return func(entry *puzzle.Entry[T], schema *JSONSchemaEntry) error {
		schema.Description = entry.Metadata.Description
		return fun(entry, schema)
	}
}

func boolJsonSchema(entry *puzzle.Entry[bool], schema *JSONSchemaEntry) error {
	schema.Type = "boolean"
	return nil
}

func durationJsonSchema(entry *puzzle.Entry[time.Duration], schema *JSONSchemaEntry) error {
	schema.Type = "string"
	schema.Pattern = "^[+-]?(?:\\d+(?:\\.\\d+)?(?:ns|us|µs|ms|s|m|h))+$"
	return nil
}

func float32JsonSchema(entry *puzzle.Entry[float32], schema *JSONSchemaEntry) error {
	schema.Type = "number"
	return nil
}

func float64JsonSchema(entry *puzzle.Entry[float64], schema *JSONSchemaEntry) error {
	schema.Type = "number"
	return nil
}

func intJsonSchema(entry *puzzle.Entry[int], schema *JSONSchemaEntry) error {
	schema.Type = "integer"
	max := uint64(int64(^UZERO >> 1))
	min := max + 1
	schema.Minimum = &MinMax{Negative: &min}
	schema.Maximum = &MinMax{Positive: &max}
	return nil
}

func int8JsonSchema(entry *puzzle.Entry[int8], schema *JSONSchemaEntry) error {
	schema.Type = "integer"
	max := uint64(1<<7 - 1)
	min := max + 1
	schema.Minimum = &MinMax{Negative: &min}
	schema.Maximum = &MinMax{Positive: &max}
	return nil
}

func int16JsonSchema(entry *puzzle.Entry[int16], schema *JSONSchemaEntry) error {
	schema.Type = "integer"
	max := uint64((1 << 15) - 1)
	min := max + 1
	schema.Minimum = &MinMax{Negative: &min}
	schema.Maximum = &MinMax{Positive: &max}
	return nil
}

func int32JsonSchema(entry *puzzle.Entry[int32], schema *JSONSchemaEntry) error {
	schema.Type = "integer"
	max := uint64(1<<31 - 1)
	min := max + 1
	schema.Minimum = &MinMax{Negative: &min}
	schema.Maximum = &MinMax{Positive: &max}
	return nil
}

func int64JsonSchema(entry *puzzle.Entry[int64], schema *JSONSchemaEntry) error {
	schema.Type = "integer"
	max := uint64(1<<63 - 1)
	min := max + 1
	schema.Minimum = &MinMax{Negative: &min}
	schema.Maximum = &MinMax{Positive: &max}
	return nil
}

func uintJsonSchema(entry *puzzle.Entry[uint], schema *JSONSchemaEntry) error {
	schema.Type = "integer"
	min := uint64(0)
	max := uint64(^UZERO)
	schema.Minimum = &MinMax{Positive: &min}
	schema.Maximum = &MinMax{Positive: &max}
	return nil
}

func uint8JsonSchema(entry *puzzle.Entry[uint8], schema *JSONSchemaEntry) error {
	schema.Type = "integer"
	min := uint64(0)
	max := uint64((1 << 8) - 1)
	schema.Minimum = &MinMax{Positive: &min}
	schema.Maximum = &MinMax{Positive: &max}
	return nil
}

func uint16JsonSchema(entry *puzzle.Entry[uint16], schema *JSONSchemaEntry) error {
	schema.Type = "integer"
	min := uint64(0)
	max := uint64((1 << 16) - 1)
	schema.Minimum = &MinMax{Positive: &min}
	schema.Maximum = &MinMax{Positive: &max}
	return nil
}

func uint32JsonSchema(entry *puzzle.Entry[uint32], schema *JSONSchemaEntry) error {
	schema.Type = "integer"
	min := uint64(0)
	max := uint64((1 << 32) - 1)
	schema.Minimum = &MinMax{Positive: &min}
	schema.Maximum = &MinMax{Positive: &max}
	return nil
}

func uint64JsonSchema(entry *puzzle.Entry[uint64], schema *JSONSchemaEntry) error {
	schema.Type = "integer"
	min := uint64(0)
	max := uint64((1 << 64) - 1)
	schema.Minimum = &MinMax{Positive: &min}
	schema.Maximum = &MinMax{Positive: &max}
	return nil
}

func stringJsonSchema(entry *puzzle.Entry[string], schema *JSONSchemaEntry) error {
	schema.Type = "string"
	return nil
}

func bytesJsonSchema(entry *puzzle.Entry[[]byte], schema *JSONSchemaEntry) error {
	schema.Type = "string"
	switch entry.Metadata.Format {
	case "base64":
		// Groups of 4 base64 characters
		// Optionally ends with 2 chars + == OR 3 chars + =
		schema.ContentEncoding = "base64"
	case "base32":
		schema.ContentEncoding = "base32"
	default: // hex
		schema.ContentEncoding = "base16"
	}
	return nil
}

func ipJsonSchema(entry *puzzle.Entry[net.IP], schema *JSONSchemaEntry) error {
	schema.Type = "string"
	schema.OneOf = []*JSONSchemaEntry{
		{Format: "ipv4"},
		{Format: "ipv6"},
	}
	return nil
}

func stringSliceJsonSchema(entry *puzzle.Entry[[]string], schema *JSONSchemaEntry) error {
	schema.Type = "array"
	schema.Items = &JSONSchemaEntry{
		Type: "string",
	}
	return nil
}
