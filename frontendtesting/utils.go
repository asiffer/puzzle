package frontendtesting

import (
	"crypto/rand"
	"fmt"
	"net"
	"time"

	"github.com/asiffer/puzzle"
	"github.com/brianvoe/gofakeit/v7"
)

func SliceEqual[T comparable](a, b []T) bool {
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

func RandomBytes(n int) []byte {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return b
}

type AllTypes struct {
	IB    bool
	B     bool
	S     string
	I     int
	I8    int8
	I16   int16
	I32   int32
	I64   int64
	U     uint
	U8    uint8
	U16   uint16
	U32   uint32
	U64   uint64
	F32   float32
	F64   float64
	D     time.Duration
	IP    net.IP
	Bytes []byte
	SS    []string
}

func RandomValues() *AllTypes {
	return &AllTypes{
		IB:    false, // this value is ignored
		B:     gofakeit.Bool(),
		S:     gofakeit.Word(),
		I:     gofakeit.Int(),
		I8:    gofakeit.Int8(),
		I16:   gofakeit.Int16(),
		I32:   gofakeit.Int32(),
		I64:   gofakeit.Int64(),
		U:     gofakeit.Uint(),
		U8:    gofakeit.Uint8(),
		U16:   gofakeit.Uint16(),
		U32:   gofakeit.Uint32(),
		U64:   gofakeit.Uint64(),
		F32:   gofakeit.Float32(),
		F64:   gofakeit.Float64(),
		D:     time.Duration(gofakeit.Int64()),
		IP:    net.ParseIP(gofakeit.IPv4Address()),
		Bytes: RandomBytes(64),
		SS:    []string{gofakeit.Word(), gofakeit.Word(), gofakeit.Word()},
	}
}

func (initial *AllTypes) Compare(other *AllTypes) error {
	if initial.IB != other.IB {
		return fmt.Errorf("expected %v, got %v", other.IB, initial.IB)
	}
	if initial.B != other.B {
		return fmt.Errorf("expected %v, got %v", other.B, initial.B)
	}
	if initial.S != other.S {
		return fmt.Errorf("expected %v, got %v", other.S, initial.S)
	}
	if initial.I != other.I {
		return fmt.Errorf("expected %v, got %v", other.I, initial.I)
	}
	if initial.I8 != other.I8 {
		return fmt.Errorf("expected %v, got %v", other.I8, initial.I8)
	}
	if initial.I16 != other.I16 {
		return fmt.Errorf("expected %v, got %v", other.I16, initial.I16)
	}
	if initial.I32 != other.I32 {
		return fmt.Errorf("expected %v, got %v", other.I32, initial.I32)
	}
	if initial.I64 != other.I64 {
		return fmt.Errorf("expected %v, got %v", other.I64, initial.I64)
	}
	if initial.U != other.U {
		return fmt.Errorf("expected %v, got %v", other.U, initial.U)
	}
	if initial.U8 != other.U8 {
		return fmt.Errorf("expected %v, got %v", other.U8, initial.U8)
	}
	if initial.U16 != other.U16 {
		return fmt.Errorf("expected %v, got %v", other.U16, initial.U16)
	}
	if initial.U32 != other.U32 {
		return fmt.Errorf("expected %v, got %v", other.U32, initial.U32)
	}
	if initial.U64 != other.U64 {
		return fmt.Errorf("expected %v, got %v", other.U64, initial.U64)
	}
	if initial.F32 != other.F32 {
		return fmt.Errorf("expected %v, got %v", other.F32, initial.F32)
	}
	if initial.F64 != other.F64 {
		return fmt.Errorf("expected %v, got %v", other.F64, initial.F64)
	}
	if initial.D != other.D {
		return fmt.Errorf("expected %v, got %v", other.D, initial.D)
	}
	if !initial.IP.Equal(other.IP) {
		return fmt.Errorf("expected %v, got %v", other.IP, initial.IP)
	}
	if !SliceEqual(initial.Bytes, other.Bytes) {
		return fmt.Errorf("expected %v, got %v", other.Bytes, initial.Bytes)
	}
	if !SliceEqual(initial.SS, other.SS) {
		return fmt.Errorf("expected %v, got %v", other.SS, initial.SS)
	}
	return nil
}

func NestedKey(key string, prefix string) string {
	if len(key) == 0 {
		return key
	}
	return fmt.Sprintf("%s%s%s", prefix, puzzle.DEFAULT_NESTING_SEPARATOR, key)
}

func RandomConfig() (*puzzle.Config, *AllTypes) {
	config := puzzle.NewConfig()
	defaultValues := RandomValues()

	puzzle.DefineVar(config, "bool", &defaultValues.B)
	puzzle.DefineVar(config, "string", &defaultValues.S)
	puzzle.DefineVar(config, "int", &defaultValues.I)
	puzzle.DefineVar(config, "int8", &defaultValues.I8)
	puzzle.DefineVar(config, "int16", &defaultValues.I16)
	puzzle.DefineVar(config, "int32", &defaultValues.I32)
	puzzle.DefineVar(config, "int64", &defaultValues.I64)
	puzzle.DefineVar(config, "uint", &defaultValues.U)
	puzzle.DefineVar(config, "uint8", &defaultValues.U8)
	puzzle.DefineVar(config, "uint16", &defaultValues.U16)
	puzzle.DefineVar(config, "uint32", &defaultValues.U32)
	puzzle.DefineVar(config, "uint64", &defaultValues.U64)
	puzzle.DefineVar(config, "float32", &defaultValues.F32)
	puzzle.DefineVar(config, "float64", &defaultValues.F64)
	puzzle.DefineVar(config, "duration", &defaultValues.D)
	puzzle.DefineVar(config, "ip", &defaultValues.IP)
	puzzle.DefineVar(config, "bytes", &defaultValues.Bytes, puzzle.WithFormat("base64"))
	puzzle.DefineVar(config, "string-slice", &defaultValues.SS)

	return config, defaultValues
}

func RandomNestedConfig() (*puzzle.Config, *AllTypes) {
	config := puzzle.NewConfig()
	defaultValues := RandomValues()

	puzzle.DefineVar(config, "bool", &defaultValues.B)
	puzzle.DefineVar(config, "string", &defaultValues.S)
	puzzle.DefineVar(config, NestedKey("int", "a"), &defaultValues.I)
	puzzle.DefineVar(config, NestedKey("int8", "a"), &defaultValues.I8)
	puzzle.DefineVar(config, NestedKey("int16", "a"), &defaultValues.I16)
	puzzle.DefineVar(config, NestedKey("int32", "b"), &defaultValues.I32)
	puzzle.DefineVar(config, NestedKey("int64", "b"), &defaultValues.I64)
	puzzle.DefineVar(config, NestedKey(NestedKey("uint", "c"), "a"), &defaultValues.U)
	puzzle.DefineVar(config, NestedKey(NestedKey("uint8", "c"), "a"), &defaultValues.U8)
	puzzle.DefineVar(config, NestedKey(NestedKey("uint16", "d"), "a"), &defaultValues.U16)
	puzzle.DefineVar(config, NestedKey(NestedKey("uint32", "c"), "b"), &defaultValues.U32)
	puzzle.DefineVar(config, NestedKey(NestedKey("uint64", "x"), "y"), &defaultValues.U64)
	puzzle.DefineVar(config, "float32", &defaultValues.F32)
	puzzle.DefineVar(config, "float64", &defaultValues.F64)
	puzzle.DefineVar(config, "duration", &defaultValues.D)
	puzzle.DefineVar(config, "ip", &defaultValues.IP)
	puzzle.DefineVar(config, "bytes", &defaultValues.Bytes, puzzle.WithFormat("base64"))
	puzzle.DefineVar(config, "string-slice", &defaultValues.SS)

	return config, defaultValues
}

func RandomConfigWithShort() (*puzzle.Config, *AllTypes) {
	config := puzzle.NewConfig()
	defaultValues := RandomValues()

	puzzle.DefineVar(config, "bool", &defaultValues.B, puzzle.WithShortFlagName("b"))
	puzzle.DefineVar(config, "string", &defaultValues.S, puzzle.WithShortFlagName("c"))
	puzzle.DefineVar(config, "int", &defaultValues.I, puzzle.WithShortFlagName("i"))
	puzzle.DefineVar(config, "int8", &defaultValues.I8, puzzle.WithShortFlagName("j"))
	puzzle.DefineVar(config, "int16", &defaultValues.I16, puzzle.WithShortFlagName("k"))
	puzzle.DefineVar(config, "int32", &defaultValues.I32, puzzle.WithShortFlagName("l"))
	puzzle.DefineVar(config, "int64", &defaultValues.I64, puzzle.WithShortFlagName("m"))
	puzzle.DefineVar(config, "uint", &defaultValues.U, puzzle.WithShortFlagName("n"))
	puzzle.DefineVar(config, "uint8", &defaultValues.U8, puzzle.WithShortFlagName("o"))
	puzzle.DefineVar(config, "uint16", &defaultValues.U16, puzzle.WithShortFlagName("p"))
	puzzle.DefineVar(config, "uint32", &defaultValues.U32, puzzle.WithShortFlagName("q"))
	puzzle.DefineVar(config, "uint64", &defaultValues.U64, puzzle.WithShortFlagName("r"))
	puzzle.DefineVar(config, "float32", &defaultValues.F32, puzzle.WithShortFlagName("s"))
	puzzle.DefineVar(config, "float64", &defaultValues.F64, puzzle.WithShortFlagName("t"))
	puzzle.DefineVar(config, "duration", &defaultValues.D, puzzle.WithShortFlagName("d"))
	puzzle.DefineVar(config, "ip", &defaultValues.IP, puzzle.WithShortFlagName("e"))
	puzzle.DefineVar(config, "bytes", &defaultValues.Bytes, puzzle.WithFormat("base64"), puzzle.WithShortFlagName("f"))
	puzzle.DefineVar(config, "string-slice", &defaultValues.SS, puzzle.WithShortFlagName("g"))

	return config, defaultValues
}
