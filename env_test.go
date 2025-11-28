package puzzle

import (
	"fmt"
	"net"
	"os"
	"strings"
	"testing"
	"time"

	"encoding/hex"

	"github.com/brianvoe/gofakeit/v7"
)

func TestEnvIgnore(t *testing.T) {
	config := NewConfig()
	initial := gofakeit.Uint64()
	x := initial
	DefineVar(config, "uint64", &x, WithoutEnv())

	if err := ReadEnv(config); err != nil {
		t.Fatalf("failed to read env: %v", err)
	}
	if x != initial {
		t.Errorf("expected %v, got %v", initial, x)
	}
}

func TestEnvNotNotDefined(t *testing.T) {
	config := NewConfig()
	initial := gofakeit.Uint64()
	x := initial
	DefineVar(config, "uint64", &x, WithEnvName(strings.ToUpper(gofakeit.Word())))

	if err := ReadEnv(config); err != nil {
		t.Fatalf("failed to read env: %v", err)
	}
	if x != initial {
		t.Errorf("expected %v, got %v", initial, x)
	}
}

// func testEnv(t *testing.T, i int) {
// 	gofakeit.Seed(i)
// 	config, initial := randomConfig()
// 	// set new values from env
// 	envValues := randomValues()
// 	var value string
// 	for entry := range config.Entries() {
// 		switch entry.GetValue().(type) {
// 		case bool:
// 			value = fmt.Sprintf("%v", envValues.b)
// 		case string:
// 			value = envValues.s
// 		case int:
// 			value = fmt.Sprintf("%v", envValues.i)
// 		case int8:
// 			value = fmt.Sprintf("%v", envValues.i8)
// 		case int16:
// 			value = fmt.Sprintf("%v", envValues.i16)
// 		case int32:
// 			value = fmt.Sprintf("%v", envValues.i32)
// 		case int64:
// 			value = fmt.Sprintf("%v", envValues.i64)
// 		case uint:
// 			value = fmt.Sprintf("%v", envValues.u)
// 		case uint8:
// 			value = fmt.Sprintf("%v", envValues.u8)
// 		case uint16:
// 			value = fmt.Sprintf("%v", envValues.u16)
// 		case uint32:
// 			value = fmt.Sprintf("%v", envValues.u32)
// 		case uint64:
// 			value = fmt.Sprintf("%v", envValues.u64)
// 		case float32:
// 			value = fmt.Sprintf("%v", envValues.f32)
// 		case float64:
// 			value = fmt.Sprintf("%v", envValues.f64)
// 		case time.Duration:
// 			value = fmt.Sprintf("%v", envValues.d)
// 		case net.IP:
// 			value = envValues.ip.String()
// 		case []string:
// 			value = strings.Join(envValues.ss, ",")
// 		case []byte:
// 			value = hex.EncodeToString(envValues.bytes)
// 		default:
// 			t.Fatalf("unsupported type %T", entry.GetValue())
// 		}
// 		os.Setenv(entry.GetMetadata().EnvName, value)
// 	}

// 	if err := ReadEnv(config); err != nil {
// 		t.Fatalf("failed to read env: %v", err)
// 	}

// 	if initial.b != envValues.b {
// 		t.Errorf("expected %v, got %v", envValues.b, initial.b)
// 	}
// 	if initial.s != envValues.s {
// 		t.Errorf("expected %v, got %v", envValues.s, initial.s)
// 	}
// 	if initial.i != envValues.i {
// 		t.Errorf("expected %v, got %v", envValues.i, initial.i)
// 	}
// 	if initial.i8 != envValues.i8 {
// 		t.Errorf("expected %v, got %v", envValues.i8, initial.i8)
// 	}
// 	if initial.i16 != envValues.i16 {
// 		t.Errorf("expected %v, got %v", envValues.i16, initial.i16)
// 	}
// 	if initial.i32 != envValues.i32 {
// 		t.Errorf("expected %v, got %v", envValues.i32, initial.i32)
// 	}
// 	if initial.i64 != envValues.i64 {
// 		t.Errorf("expected %v, got %v", envValues.i64, initial.i64)
// 	}
// 	if initial.u != envValues.u {
// 		t.Errorf("expected %v, got %v", envValues.u, initial.u)
// 	}
// 	if initial.u8 != envValues.u8 {
// 		t.Errorf("expected %v, got %v", envValues.u8, initial.u8)
// 	}
// 	if initial.u16 != envValues.u16 {
// 		t.Errorf("expected %v, got %v", envValues.u16, initial.u16)
// 	}
// 	if initial.u32 != envValues.u32 {
// 		t.Errorf("expected %v, got %v", envValues.u32, initial.u32)
// 	}
// 	if initial.u64 != envValues.u64 {
// 		t.Errorf("expected %v, got %v", envValues.u64, initial.u64)
// 	}
// 	if initial.f32 != envValues.f32 {
// 		t.Errorf("expected %v, got %v", envValues.f32, initial.f32)
// 	}
// 	if initial.f64 != envValues.f64 {
// 		t.Errorf("expected %v, got %v", envValues.f64, initial.f64)
// 	}
// 	if initial.d != envValues.d {
// 		t.Errorf("expected %v, got %v", envValues.d, initial.d)
// 	}
// 	if !initial.ip.Equal(envValues.ip) {
// 		t.Errorf("expected %v, got %v", envValues.ip, initial.ip)
// 	}
// 	if !sliceEqualFactory(envValues.bytes)(initial.bytes) {
// 		t.Errorf("expected %v, got %v", envValues.ip, initial.ip)
// 	}
// }

func testGenericRead(t *testing.T, setter func(entry EntryInterface, value string), reader func(c *Config) error, i int) {
	gofakeit.Seed(i)
	config, initial := randomConfig()
	// set new values from env
	readValues := randomValues()
	var value string
	for entry := range config.Entries() {
		switch entry.GetValue().(type) {
		case bool:
			value = fmt.Sprintf("%v", readValues.b)
		case string:
			value = readValues.s
		case int:
			value = fmt.Sprintf("%v", readValues.i)
		case int8:
			value = fmt.Sprintf("%v", readValues.i8)
		case int16:
			value = fmt.Sprintf("%v", readValues.i16)
		case int32:
			value = fmt.Sprintf("%v", readValues.i32)
		case int64:
			value = fmt.Sprintf("%v", readValues.i64)
		case uint:
			value = fmt.Sprintf("%v", readValues.u)
		case uint8:
			value = fmt.Sprintf("%v", readValues.u8)
		case uint16:
			value = fmt.Sprintf("%v", readValues.u16)
		case uint32:
			value = fmt.Sprintf("%v", readValues.u32)
		case uint64:
			value = fmt.Sprintf("%v", readValues.u64)
		case float32:
			value = fmt.Sprintf("%v", readValues.f32)
		case float64:
			value = fmt.Sprintf("%v", readValues.f64)
		case time.Duration:
			value = fmt.Sprintf("%v", readValues.d)
		case net.IP:
			value = readValues.ip.String()
		case []string:
			value = strings.Join(readValues.ss, ",")
		case []byte:
			value = hex.EncodeToString(readValues.bytes)
		default:
			t.Fatalf("unsupported type %T", entry.GetValue())
		}
		setter(entry, value)
	}

	if err := reader(config); err != nil {
		t.Fatalf("failed to read: %v", err)
	}

	if initial.b != readValues.b {
		t.Errorf("expected %v, got %v", readValues.b, initial.b)
	}
	if initial.s != readValues.s {
		t.Errorf("expected %v, got %v", readValues.s, initial.s)
	}
	if initial.i != readValues.i {
		t.Errorf("expected %v, got %v", readValues.i, initial.i)
	}
	if initial.i8 != readValues.i8 {
		t.Errorf("expected %v, got %v", readValues.i8, initial.i8)
	}
	if initial.i16 != readValues.i16 {
		t.Errorf("expected %v, got %v", readValues.i16, initial.i16)
	}
	if initial.i32 != readValues.i32 {
		t.Errorf("expected %v, got %v", readValues.i32, initial.i32)
	}
	if initial.i64 != readValues.i64 {
		t.Errorf("expected %v, got %v", readValues.i64, initial.i64)
	}
	if initial.u != readValues.u {
		t.Errorf("expected %v, got %v", readValues.u, initial.u)
	}
	if initial.u8 != readValues.u8 {
		t.Errorf("expected %v, got %v", readValues.u8, initial.u8)
	}
	if initial.u16 != readValues.u16 {
		t.Errorf("expected %v, got %v", readValues.u16, initial.u16)
	}
	if initial.u32 != readValues.u32 {
		t.Errorf("expected %v, got %v", readValues.u32, initial.u32)
	}
	if initial.u64 != readValues.u64 {
		t.Errorf("expected %v, got %v", readValues.u64, initial.u64)
	}
	if initial.f32 != readValues.f32 {
		t.Errorf("expected %v, got %v", readValues.f32, initial.f32)
	}
	if initial.f64 != readValues.f64 {
		t.Errorf("expected %v, got %v", readValues.f64, initial.f64)
	}
	if initial.d != readValues.d {
		t.Errorf("expected %v, got %v", readValues.d, initial.d)
	}
	if !initial.ip.Equal(readValues.ip) {
		t.Errorf("expected %v, got %v", readValues.ip, initial.ip)
	}
	if !sliceEqualFactory(readValues.bytes)(initial.bytes) {
		t.Errorf("expected %v, got %v", readValues.ip, initial.ip)
	}
}

func testEnv(t *testing.T, i int) {
	setter := func(entry EntryInterface, value string) {
		os.Setenv(entry.GetMetadata().EnvName, value)
	}
	reader := func(c *Config) error {
		return ReadEnv(c)
	}
	testGenericRead(t, setter, reader, i)
}

func FuzzEnv(f *testing.F) {
	for i := 0; i < FUZZ_SIZE; i++ {
		f.Add(i)
	}

	f.Fuzz(testEnv)
}
