package puzzle

import (
	"encoding/base32"
	"encoding/base64"
	"encoding/csv"
	"encoding/hex"
	"fmt"
	"io"
	"net"
	"strings"
	"time"
)

// Entry is a structure that represents an item in the configuration
type Entry[T any] struct {
	Metadata  *EntryMetadata
	Key       string
	ValueP    *T // to bind the Value if not stored by the entry
	Value     T
	converter Converter[T] // pivot structure for frontends
}

func NewEntry[T any](Key string) *Entry[T] {
	e := Entry[T]{Key: Key}
	// by default we bind the to the local storage
	e.ValueP = &e.Value
	e.autoMetadata()
	e.wire()
	return &e
}

func (e *Entry[T]) autoMetadata() {
	e.Metadata = newMetadataFromEntry(e.Key)
	var t T
	if _, ok := any(t).(bool); ok {
		e.Metadata.IsBool = true
	}
}

// Wire performs all the plumbing
func (e *Entry[T]) wire() {
	switch z := any(e).(type) {
	case *Entry[bool]:
		z.converter = BoolConverter
	case *Entry[time.Duration]:
		z.converter = DurationConverter
	case *Entry[float32]:
		z.converter = Float32Converter
	case *Entry[float64]:
		z.converter = Float64Converter
	case *Entry[int]:
		z.converter = IntConverter
	case *Entry[int8]:
		z.converter = Int8Converter
	case *Entry[int16]:
		z.converter = Int16Converter
	case *Entry[int32]:
		z.converter = Int32Converter
	case *Entry[int64]:
		z.converter = Int64Converter
	case *Entry[string]:
		z.converter = StringConverter
	case *Entry[uint]:
		z.converter = UintConverter
	case *Entry[uint8]:
		z.converter = Uint8Converter
	case *Entry[uint16]:
		z.converter = Uint16Converter
	case *Entry[uint32]:
		z.converter = Uint32Converter
	case *Entry[uint64]:
		z.converter = Uint64Converter
	case *Entry[[]byte]:
		z.converter = BytesConverter
	case *Entry[[]string]:
		z.converter = StringSliceConverter
	case *Entry[net.IP]:
		z.converter = IPConverter
	default:
		panic(fmt.Sprintf("unsupported type %T", e.Value))
	}
}

// EntryInterface is the object that unifies all the provided entries
// (that have different types)
type EntryInterface interface {
	GetKey() string
	GetValue() interface{}
	GetMetadata() *EntryMetadata

	Set(string) error
	String() string
	Convert(frontend Frontend, args ...any) error
}

// Join a slice of strings into a single string using the entry's slice separator
func join(s []string, entry EntryInterface) string {
	var out strings.Builder
	w := csv.NewWriter(&out)
	w.Comma = entry.GetMetadata().SliceSeparator
	w.Write(s)
	w.Flush()
	return out.String()
}

// GetKey returns the key of the entry
func (e *Entry[T]) GetKey() string {
	return e.Key
}

// GetValue returns the value of the entry as an interface{}.
func (e *Entry[T]) GetValue() interface{} {
	return *e.ValueP
}

// GetMetadata returns the metadata of the entry
func (e *Entry[T]) GetMetadata() *EntryMetadata {
	return e.Metadata
}

// Method to be compatible with flag.Value interface (and spf13/pflag.Value interface)
func (e *Entry[T]) String() string {
	// possible issue with flag package. See https://pkg.go.dev/flag#Value
	if e == nil || e.ValueP == nil {
		return ""
	}
	switch v := any(*e.ValueP).(type) {
	case time.Duration:
		return v.String()
	case net.IP:
		return v.String()
	case []string:
		var out strings.Builder
		w := e.csvWriter(&out)
		w.Write(v)
		w.Flush()
		return out.String()
	case []byte:
		switch e.Metadata.Format {
		case "base32":
			return base32.StdEncoding.EncodeToString(v)
		case "base64":
			return base64.StdEncoding.EncodeToString(v)
		default: // default to hex
			return hex.EncodeToString(v)
		}
	default:
		return fmt.Sprintf("%v", *e.ValueP)
	}
}

// Convert converts the entry value using the provided frontend and arguments.
// It updates the value of the entry.
func (e *Entry[T]) Convert(frontend Frontend, args ...any) error {
	return e.converter.Convert(frontend, e, args...)
}

// Method to be compatible with flag.Value interface (and spf13/pflag.Value interface)
func (e *Entry[T]) Set(value string) error {
	return e.Convert(StringFrontend, value)
}

// Method to be compatible with spf13/pflag.Value interface
func (e *Entry[T]) Type() string {
	return fmt.Sprintf("%T", e.Value)
}

// Method to be compatible with urfave/cli.Value interface
func (e *Entry[T]) Get() interface{} {
	return e.GetValue()
}

func (e *Entry[T]) csvReader(r io.Reader) *csv.Reader {
	reader := csv.NewReader(r)
	reader.Comma = e.Metadata.SliceSeparator
	return reader
}

func (e *Entry[T]) csvWriter(w io.Writer) *csv.Writer {
	writer := csv.NewWriter(w)
	writer.Comma = e.Metadata.SliceSeparator
	return writer
}

// For flag package compatibility
// see https://pkg.go.dev/flag#Value
func (e *Entry[T]) IsBoolFlag() bool {
	return e.Metadata.IsBool
}
