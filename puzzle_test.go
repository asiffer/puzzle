package puzzle

import (
	"fmt"
	"net"
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

const FUZZ_SIZE = 200

type EntryModifier[T any] func(e *Entry[T])

type allTypes struct {
	b     bool
	s     string
	i     int
	i8    int8
	i16   int16
	i32   int32
	i64   int64
	u     uint
	u8    uint8
	u16   uint16
	u32   uint32
	u64   uint64
	f32   float32
	f64   float64
	d     time.Duration
	ip    net.IP
	bytes []byte
	ss    []string
}

func randomValues() *allTypes {
	b, err := randomBytes(64)
	if err != nil {
		panic(err)
	}
	return &allTypes{
		b:     gofakeit.Bool(),
		s:     gofakeit.Word(),
		i:     gofakeit.Int(),
		i8:    gofakeit.Int8(),
		i16:   gofakeit.Int16(),
		i32:   gofakeit.Int32(),
		i64:   gofakeit.Int64(),
		u:     gofakeit.Uint(),
		u8:    gofakeit.Uint8(),
		u16:   gofakeit.Uint16(),
		u32:   gofakeit.Uint32(),
		u64:   gofakeit.Uint64(),
		f32:   gofakeit.Float32(),
		f64:   gofakeit.Float64(),
		d:     time.Duration(gofakeit.Int64()),
		ip:    net.ParseIP(gofakeit.IPv4Address()),
		bytes: b,
		ss:    []string{gofakeit.Word(), gofakeit.Word(), gofakeit.Word()},
	}
}

func randomConfig() (*Config, *allTypes) {
	config := NewConfig()
	defaultValues := randomValues()

	DefineVar(config, "bool", &defaultValues.b)
	DefineVar(config, "string", &defaultValues.s)
	DefineVar(config, "int", &defaultValues.i)
	DefineVar(config, "int8", &defaultValues.i8)
	DefineVar(config, "int16", &defaultValues.i16)
	DefineVar(config, "int32", &defaultValues.i32)
	DefineVar(config, "int64", &defaultValues.i64)
	DefineVar(config, "uint", &defaultValues.u)
	DefineVar(config, "uint8", &defaultValues.u8)
	DefineVar(config, "uint16", &defaultValues.u16)
	DefineVar(config, "uint32", &defaultValues.u32)
	DefineVar(config, "uint64", &defaultValues.u64)
	DefineVar(config, "float32", &defaultValues.f32)
	DefineVar(config, "float64", &defaultValues.f64)
	DefineVar(config, "duration", &defaultValues.d)
	DefineVar(config, "ip", &defaultValues.ip)
	DefineVar(config, "bytes", &defaultValues.bytes)
	DefineVar(config, "string-slice", &defaultValues.ss)

	return config, defaultValues
}

func sliceEqualFactory[T comparable](b []T) func(a []T) bool {
	return func(a []T) bool {
		if len(a) != len(b) {
			return false
		}
		for i := range a {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}
}

func testConverter[T comparable](converter Converter[T], stringValue string, expected T, modifiers ...EntryModifier[T]) error {
	return testConverterAny(converter, stringValue, expected, func(v T) bool {
		return v == expected
	}, modifiers...)
}

func testConverterAny[T any](converter Converter[T], stringValue string, expected T, check func(v T) bool, modifiers ...EntryModifier[T]) error {
	e := NewEntry[T](gofakeit.LetterN(8))
	for _, mod := range modifiers {
		mod(e)
	}

	if err := converter.Convert(StringFrontend, e, stringValue); err != nil {
		return fmt.Errorf("Error converting %s to %T: %v", stringValue, e.Value, err)
	}
	if !check(e.Value) {
		fmt.Println(check(e.Value))
		return fmt.Errorf("Expected %v, got %v (type: %T)", expected, e.Value, e.Value)
	}
	return nil
}
