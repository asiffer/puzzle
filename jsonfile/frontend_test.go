package jsonfile

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/asiffer/puzzle"
	"github.com/asiffer/puzzle/frontendtesting"
	"github.com/brianvoe/gofakeit/v7"
)

// func toJSON(config *puzzle.Config) ([]byte, error) {
// 	tmp := make(map[string]interface{})
// 	for entry := range config.Entries() {
// 		k := entry.GetKey()
// 		if strings.Contains(k, config.NestingSeparator) {
// 			keys := strings.Split(k, config.NestingSeparator)
// 			tmp0 := tmp
// 			for i := 0; i < len(keys)-1; i++ {
// 				if _, ok := tmp0[keys[i]]; !ok {
// 					tmp0[keys[i]] = make(map[string]interface{})
// 				}
// 				tmp0 = tmp0[keys[i]].(map[string]interface{})
// 			}
// 			k = keys[len(keys)-1]
// 			tmp0[k] = entry.GetValue()
// 		} else {
// 			tmp[k] = entry.GetValue()
// 		}

// 		switch v := tmp[k].(type) {
// 		case net.IP:
// 			tmp[k] = v.String()
// 		case time.Duration:
// 			tmp[k] = v.String()
// 		case []byte:
// 			tmp[k] = base64.StdEncoding.EncodeToString(v)
// 		}
// 	}
// 	return json.Marshal(tmp)
// }

func randomConfig() (*puzzle.Config, *frontendtesting.AllTypes) {
	config := puzzle.NewConfig()
	defaultValues := frontendtesting.RandomValues()

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

func nestedKey(key string, prefix string) string {
	if len(key) == 0 {
		return key
	}
	return fmt.Sprintf("%s%s%s", prefix, puzzle.DEFAULT_NESTING_SEPARATOR, key)
}

func randomNestedConfig() (*puzzle.Config, *frontendtesting.AllTypes) {
	config := puzzle.NewConfig()
	defaultValues := frontendtesting.RandomValues()

	puzzle.DefineVar(config, "bool", &defaultValues.B)
	puzzle.DefineVar(config, "string", &defaultValues.S)
	puzzle.DefineVar(config, nestedKey("int", "a"), &defaultValues.I)
	puzzle.DefineVar(config, nestedKey("int8", "a"), &defaultValues.I8)
	puzzle.DefineVar(config, nestedKey("int16", "a"), &defaultValues.I16)
	puzzle.DefineVar(config, nestedKey("int32", "b"), &defaultValues.I32)
	puzzle.DefineVar(config, nestedKey("int64", "b"), &defaultValues.I64)
	puzzle.DefineVar(config, nestedKey(nestedKey("uint", "c"), "a"), &defaultValues.U)
	puzzle.DefineVar(config, nestedKey(nestedKey("uint8", "c"), "a"), &defaultValues.U8)
	puzzle.DefineVar(config, nestedKey(nestedKey("uint16", "d"), "a"), &defaultValues.U16)
	puzzle.DefineVar(config, nestedKey(nestedKey("uint32", "c"), "b"), &defaultValues.U32)
	puzzle.DefineVar(config, nestedKey(nestedKey("uint64", "x"), "y"), &defaultValues.U64)
	puzzle.DefineVar(config, "float32", &defaultValues.F32)
	puzzle.DefineVar(config, "float64", &defaultValues.F64)
	puzzle.DefineVar(config, "duration", &defaultValues.D)
	puzzle.DefineVar(config, "ip", &defaultValues.IP)
	puzzle.DefineVar(config, "bytes", &defaultValues.Bytes, puzzle.WithFormat("base64"))
	puzzle.DefineVar(config, "string-slice", &defaultValues.SS)

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

func testReadJsonRaw(config *puzzle.Config, initial *frontendtesting.AllTypes, raw []byte, jsonValues *frontendtesting.AllTypes, t *testing.T) {
	t.Log(string(raw))
	if err := ReadJSONRaw(config, raw); err != nil {
		t.Fatalf("failed to read json: %v", err)
	}
	if err := initial.Compare(jsonValues); err != nil {
		t.Error(err)
	}
}

func TestReadJSONRaw(t *testing.T) {
	config, initial := randomConfig()
	config2, jsonValues := randomConfig()
	raw, err := ToJSON(config2)
	if err != nil {
		t.Fatalf("failed to marshal json: %v", err)
	}
	testReadJsonRaw(config, initial, raw, jsonValues, t)
}

func TestReadJSONRawNested(t *testing.T) {
	config, initial := randomNestedConfig()
	config2, jsonValues := randomNestedConfig()
	raw, err := ToJSON(config2)
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
	raw, err := ToJSON(config2)
	t.Log(string(raw))
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

	if err := initial.Compare(jsonValues); err != nil {
		t.Error(err)
	}
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
