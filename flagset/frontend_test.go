package flagset

import (
	"crypto/rand"
	"encoding/base64"
	"flag"
	"fmt"
	"net"
	"strings"
	"testing"
	"time"

	"github.com/asiffer/puzzle"
	"github.com/brianvoe/gofakeit/v7"
)

func toFlags(c *puzzle.Config) []string {
	out := make([]string, 0)
	for e := range c.Entries() {
		fn := e.GetMetadata().FlagName
		if fn == "" {
			continue
		}
		value := e.String()
		switch v := e.GetValue().(type) {
		case net.IP:
			value = v.String()
		case time.Duration:
			value = v.String()
		case []byte:
			value = base64.StdEncoding.EncodeToString(v)
		case []string:
			value = strings.Join(v, ",")
		}
		out = append(out, fmt.Sprintf("-%s=%v", fn, value))
	}
	return out
}

func randomBytes(n int) []byte {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return b
}

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

func sliceEqual[T comparable](a, b []T) bool {
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
func compare(initial *allTypes, jsonValues *allTypes, t *testing.T) {
	if initial.b != jsonValues.b {
		t.Errorf("expected %v, got %v", jsonValues.b, initial.b)
	}
	if initial.s != jsonValues.s {
		t.Errorf("expected %v, got %v", jsonValues.s, initial.s)
	}
	if initial.i != jsonValues.i {
		t.Errorf("expected %v, got %v", jsonValues.i, initial.i)
	}
	if initial.i8 != jsonValues.i8 {
		t.Errorf("expected %v, got %v", jsonValues.i8, initial.i8)
	}
	if initial.i16 != jsonValues.i16 {
		t.Errorf("expected %v, got %v", jsonValues.i16, initial.i16)
	}
	if initial.i32 != jsonValues.i32 {
		t.Errorf("expected %v, got %v", jsonValues.i32, initial.i32)
	}
	if initial.i64 != jsonValues.i64 {
		t.Errorf("expected %v, got %v", jsonValues.i64, initial.i64)
	}
	if initial.u != jsonValues.u {
		t.Errorf("expected %v, got %v", jsonValues.u, initial.u)
	}
	if initial.u8 != jsonValues.u8 {
		t.Errorf("expected %v, got %v", jsonValues.u8, initial.u8)
	}
	if initial.u16 != jsonValues.u16 {
		t.Errorf("expected %v, got %v", jsonValues.u16, initial.u16)
	}
	if initial.u32 != jsonValues.u32 {
		t.Errorf("expected %v, got %v", jsonValues.u32, initial.u32)
	}
	if initial.u64 != jsonValues.u64 {
		t.Errorf("expected %v, got %v", jsonValues.u64, initial.u64)
	}
	if initial.f32 != jsonValues.f32 {
		t.Errorf("expected %v, got %v", jsonValues.f32, initial.f32)
	}
	if initial.f64 != jsonValues.f64 {
		t.Errorf("expected %v, got %v", jsonValues.f64, initial.f64)
	}
	if initial.d != jsonValues.d {
		t.Errorf("expected %v, got %v", jsonValues.d, initial.d)
	}
	if !initial.ip.Equal(jsonValues.ip) {
		t.Errorf("expected %v, got %v", jsonValues.ip, initial.ip)
	}
	if !sliceEqual(initial.bytes, jsonValues.bytes) {
		t.Errorf("expected %v, got %v", jsonValues.bytes, initial.bytes)
	}
	if !sliceEqual(initial.ss, jsonValues.ss) {
		t.Errorf("expected %v, got %v", jsonValues.ss, initial.ss)
	}
}

func randomValues() *allTypes {
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
		bytes: randomBytes(64),
		ss:    []string{gofakeit.Word(), gofakeit.Word(), gofakeit.Word()},
	}
}

func randomConfig() (*puzzle.Config, *allTypes) {
	config := puzzle.NewConfig()
	defaultValues := randomValues()

	puzzle.DefineVar(config, "bool", &defaultValues.b)
	puzzle.DefineVar(config, "string", &defaultValues.s)
	puzzle.DefineVar(config, "int", &defaultValues.i)
	puzzle.DefineVar(config, "int8", &defaultValues.i8)
	puzzle.DefineVar(config, "int16", &defaultValues.i16)
	puzzle.DefineVar(config, "int32", &defaultValues.i32)
	puzzle.DefineVar(config, "int64", &defaultValues.i64)
	puzzle.DefineVar(config, "uint", &defaultValues.u)
	puzzle.DefineVar(config, "uint8", &defaultValues.u8)
	puzzle.DefineVar(config, "uint16", &defaultValues.u16)
	puzzle.DefineVar(config, "uint32", &defaultValues.u32)
	puzzle.DefineVar(config, "uint64", &defaultValues.u64)
	puzzle.DefineVar(config, "float32", &defaultValues.f32)
	puzzle.DefineVar(config, "float64", &defaultValues.f64)
	puzzle.DefineVar(config, "duration", &defaultValues.d)
	puzzle.DefineVar(config, "ip", &defaultValues.ip)
	puzzle.DefineVar(config, "bytes", &defaultValues.bytes, puzzle.WithFormat("base64"))
	puzzle.DefineVar(config, "string-slice", &defaultValues.ss)

	return config, defaultValues
}

func testBuild(t *testing.T, i int) {
	gofakeit.Seed(i)
	config, initial := randomConfig()
	config2, values := randomConfig()
	toFlags(config2)

	fs, err := Build(config, "test", flag.PanicOnError)
	if err != nil {
		t.Fatalf("error building flagset: %v", err)
	}
	if err := fs.Parse(toFlags(config2)); err != nil {
		t.Fatalf("error parsing flags: %v", err)
	}

	compare(initial, values, t)
}

func FuzzBuild(f *testing.F) {
	for i := range 200 {
		f.Add(i)
	}
	f.Fuzz(testBuild)
}
