package jsonfile

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"path"
	"strings"
	"testing"
	"time"

	"github.com/asiffer/puzzle"
	"github.com/brianvoe/gofakeit/v7"
)

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

func randomBytes(n int) []byte {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return b
}

func toJSON(config *puzzle.Config) ([]byte, error) {
	tmp := make(map[string]interface{})
	for entry := range config.Entries() {
		k := entry.GetKey()
		if strings.Contains(k, config.NestingSeparator) {
			keys := strings.Split(k, config.NestingSeparator)
			tmp0 := tmp
			for i := 0; i < len(keys)-1; i++ {
				if _, ok := tmp0[keys[i]]; !ok {
					tmp0[keys[i]] = make(map[string]interface{})
				}
				tmp0 = tmp0[keys[i]].(map[string]interface{})
			}
			k = keys[len(keys)-1]
			tmp0[k] = entry.GetValue()
		} else {
			tmp[k] = entry.GetValue()
		}

		switch v := tmp[k].(type) {
		case net.IP:
			tmp[k] = v.String()
		case time.Duration:
			tmp[k] = v.String()
		case []byte:
			tmp[k] = base64.StdEncoding.EncodeToString(v)
		}
	}
	return json.Marshal(tmp)
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

func nestedKey(key string, prefix string) string {
	if len(key) == 0 {
		return key
	}
	return fmt.Sprintf("%s%s%s", prefix, puzzle.DEFAULT_NESTING_SEPARATOR, key)
}

func randomNestedConfig() (*puzzle.Config, *allTypes) {
	config := puzzle.NewConfig()
	defaultValues := randomValues()

	puzzle.DefineVar(config, "bool", &defaultValues.b)
	puzzle.DefineVar(config, "string", &defaultValues.s)
	puzzle.DefineVar(config, nestedKey("int", "a"), &defaultValues.i)
	puzzle.DefineVar(config, nestedKey("int8", "a"), &defaultValues.i8)
	puzzle.DefineVar(config, nestedKey("int16", "a"), &defaultValues.i16)
	puzzle.DefineVar(config, nestedKey("int32", "b"), &defaultValues.i32)
	puzzle.DefineVar(config, nestedKey("int64", "b"), &defaultValues.i64)
	puzzle.DefineVar(config, nestedKey(nestedKey("uint", "c"), "a"), &defaultValues.u)
	puzzle.DefineVar(config, nestedKey(nestedKey("uint8", "c"), "a"), &defaultValues.u8)
	puzzle.DefineVar(config, nestedKey(nestedKey("uint16", "d"), "a"), &defaultValues.u16)
	puzzle.DefineVar(config, nestedKey(nestedKey("uint32", "c"), "b"), &defaultValues.u32)
	puzzle.DefineVar(config, nestedKey(nestedKey("uint64", "x"), "y"), &defaultValues.u64)
	puzzle.DefineVar(config, "float32", &defaultValues.f32)
	puzzle.DefineVar(config, "float64", &defaultValues.f64)
	puzzle.DefineVar(config, "duration", &defaultValues.d)
	puzzle.DefineVar(config, "ip", &defaultValues.ip)
	puzzle.DefineVar(config, "bytes", &defaultValues.bytes, puzzle.WithFormat("base64"))
	puzzle.DefineVar(config, "string-slice", &defaultValues.ss)

	return config, defaultValues
}

func TestReadJSONRawBadKey(t *testing.T) {
	config, _ := randomConfig()
	raw := []byte(`{"error":"bad key"}`)
	t.Log(string(raw))
	if err := ReadJSONRaw(config, raw); err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestReadJSONRawBadKey2(t *testing.T) {
	config, _ := randomConfig()
	raw := []byte(`{"error": ["bad key"]}`)
	t.Log(string(raw))
	if err := ReadJSONRaw(config, raw); err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestReadJSONRawBadJSON(t *testing.T) {
	config, _ := randomConfig()
	raw := []byte(`BAD JSON`)
	t.Log(string(raw))
	if err := ReadJSONRaw(config, raw); err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestReadJSONRawType(t *testing.T) {
	config, _ := randomConfig()
	raw := []byte(`{"bool":"an invalid boolean"}`)
	t.Log(string(raw))
	if err := ReadJSONRaw(config, raw); err == nil {
		t.Fatal("expected error, got nil")
	}
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

func testReadJsonRaw(config *puzzle.Config, initial *allTypes, raw []byte, jsonValues *allTypes, t *testing.T) {
	t.Log(string(raw))
	if err := ReadJSONRaw(config, raw); err != nil {
		t.Fatalf("failed to read json: %v", err)
	}
	compare(initial, jsonValues, t)
}

func TestReadJSONRaw(t *testing.T) {
	config, initial := randomConfig()
	config2, jsonValues := randomConfig()
	raw, err := toJSON(config2)
	if err != nil {
		t.Fatalf("failed to marshal json: %v", err)
	}
	testReadJsonRaw(config, initial, raw, jsonValues, t)
}

func TestReadJSONRawNested(t *testing.T) {
	config, initial := randomNestedConfig()
	config2, jsonValues := randomNestedConfig()
	raw, err := toJSON(config2)
	if err != nil {
		t.Fatalf("failed to marshal json: %v", err)
	}
	testReadJsonRaw(config, initial, raw, jsonValues, t)
}

func TestReadJSON(t *testing.T) {
	config, initial := randomNestedConfig()

	dir := t.TempDir()
	filename := gofakeit.LetterN(8) + ".json"
	path := path.Join(dir, filename)
	if err := puzzle.DefineConfigFile(config, "config", []string{path}); err != nil {
		t.Fatalf("failed to define config file: %v", err)
	}
	config2, jsonValues := randomNestedConfig()
	raw, err := toJSON(config2)
	if err != nil {
		t.Fatalf("failed to marshal json: %v", err)
	}
	// write raw to path
	if err := os.WriteFile(path, raw, 0644); err != nil {
		t.Fatalf("failed to write file: %v", err)
	}

	if err := ReadJSON(config); err != nil {
		t.Fatalf("failed to read json: %v", err)
	}
	compare(initial, jsonValues, t)
}

func TestReadJSONNoFile(t *testing.T) {
	config, _ := randomNestedConfig()

	dir := t.TempDir()
	filename := gofakeit.LetterN(8) + ".json"
	path := path.Join(dir, filename)
	if err := puzzle.DefineConfigFile(config, "config", []string{path}); err != nil {
		t.Fatalf("failed to define config file: %v", err)
	}

	if err := ReadJSON(config); err == nil {
		t.Fatalf("An error was expected since the config file does not exist: %v", err)
	}
}

func TestReadJSONBadFile(t *testing.T) {
	config, _ := randomNestedConfig()

	dir := t.TempDir()
	filename := gofakeit.LetterN(8) + ".json"
	path := path.Join(dir, filename)
	if err := puzzle.DefineConfigFile(config, "config", []string{path}); err != nil {
		t.Fatalf("failed to define config file: %v", err)
	}

	if err := os.WriteFile(path, []byte(`BAD JSON`), 0644); err != nil {
		t.Fatalf("failed to write file: %v", err)
	}

	if err := ReadJSON(config); err == nil {
		t.Fatalf("An error was expected since the config file does not exist: %v", err)
	}
}

func TestReadJSONNoConfig(t *testing.T) {
	config, _ := randomNestedConfig()
	if err := puzzle.DefineConfigFile(config, "config", []string{}); err != nil {
		t.Fatalf("failed to define config file: %v", err)
	}
	if err := ReadJSON(config); err == nil {
		t.Fatalf("An error was expected since the config file does not exist: %v", err)
	}
}
