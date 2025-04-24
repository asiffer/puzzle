package puzzle

import (
	"fmt"
	"net"
	"time"
)

type Entry[T any] struct {
	Metadata  *EntryMetadata
	Key       string
	ValueP    *T // to bind the Value if not stored by the entry
	Value     T
	converter Converter[T]
}

func newEntry[T any](Key string) *Entry[T] {
	e := Entry[T]{Key: Key}
	// by default we bind the to the local storage
	e.ValueP = &e.Value
	e.autoMetadata()
	e.wire()
	return &e
}

func (e *Entry[T]) autoMetadata() {
	e.Metadata = newMetadataFromEntry(e.Key)
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

type EntryInterface interface {
	GetKey() string
	GetValue() interface{}
	GetMetadata() *EntryMetadata

	Set(string) error
	String() string
	Convert(frontend Frontend, args ...any) error
}

func (e *Entry[T]) GetKey() string {
	return e.Key
}

func (e *Entry[T]) GetValue() interface{} {
	return *e.ValueP
}

func (e *Entry[T]) GetMetadata() *EntryMetadata {
	return e.Metadata
}

// Method to be compatible with flag.Value interface (and spf13/pflag.Value interface)
func (e *Entry[T]) String() string {
	return fmt.Sprintf("%v", *e.ValueP)
}

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
