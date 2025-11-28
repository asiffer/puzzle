package jsonfile

import (
	"os"
	"path"
	"testing"

	"github.com/asiffer/puzzle"
	"github.com/asiffer/puzzle/frontendtesting"
	"github.com/brianvoe/gofakeit/v7"
)

func TestReadJSONRawBadKey(t *testing.T) {
	config, _ := frontendtesting.RandomConfig()
	raw := []byte(`{"error":"bad key"}`)
	t.Log(string(raw))
	if err := ReadJSONRaw(config, raw); err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestReadJSONRawBadKey2(t *testing.T) {
	config, _ := frontendtesting.RandomConfig()
	raw := []byte(`{"error": ["bad key"]}`)
	t.Log(string(raw))
	if err := ReadJSONRaw(config, raw); err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestReadJSONRawBadJSON(t *testing.T) {
	config, _ := frontendtesting.RandomConfig()
	raw := []byte(`BAD JSON`)
	t.Log(string(raw))
	if err := ReadJSONRaw(config, raw); err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestReadJSONRawType(t *testing.T) {
	config, _ := frontendtesting.RandomConfig()
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
	config, initial := frontendtesting.RandomConfig()
	config2, jsonValues := frontendtesting.RandomConfig()
	raw, err := ToJSON(config2)
	if err != nil {
		t.Fatalf("failed to marshal json: %v", err)
	}
	testReadJsonRaw(config, initial, raw, jsonValues, t)
}

func TestReadJSONRawNested(t *testing.T) {
	config, initial := frontendtesting.RandomNestedConfig()
	config2, jsonValues := frontendtesting.RandomNestedConfig()
	raw, err := ToJSON(config2)
	if err != nil {
		t.Fatalf("failed to marshal json: %v", err)
	}
	testReadJsonRaw(config, initial, raw, jsonValues, t)
}

func TestReadJSON(t *testing.T) {
	config, initial := frontendtesting.RandomNestedConfig()

	dir := t.TempDir()
	filename := gofakeit.LetterN(8) + ".json"
	path := path.Join(dir, filename)
	if err := puzzle.DefineConfigFile(config, "config", []string{path}); err != nil {
		t.Fatalf("failed to define config file: %v", err)
	}
	config2, jsonValues := frontendtesting.RandomNestedConfig()
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
	config, _ := frontendtesting.RandomNestedConfig()

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
	config, _ := frontendtesting.RandomNestedConfig()

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
	config, _ := frontendtesting.RandomNestedConfig()
	if err := puzzle.DefineConfigFile(config, "config", []string{}); err != nil {
		t.Fatalf("failed to define config file: %v", err)
	}
	if err := ReadJSON(config); err == nil {
		t.Fatalf("An error was expected since the config file does not exist: %v", err)
	}
}
