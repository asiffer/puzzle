package jsonfile

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/asiffer/puzzle"
)

func recursiveJSON(config *puzzle.Config, data interface{}, prefix string, out map[string]string) error {
	switch val := data.(type) {
	case map[string]interface{}:
		for k, v := range val {
			key := k
			if prefix != "" {
				key = prefix + config.NestingSeparator + k
			}
			if err := recursiveJSON(config, v, key, out); err != nil {
				return err
			}
		}
	case []interface{}:
		entry, exists := config.GetEntry(prefix)
		if !exists {
			return fmt.Errorf("key %s is not accepted by the configuration", prefix)
		}

		ssep := entry.GetMetadata().SliceSeparator
		tmp := ""
		for _, v := range val {
			tmp += fmt.Sprintf("%v%s", v, ssep)
		}
		out[prefix] = strings.TrimSuffix(tmp, ssep)
		return nil
	default:
		_, exists := config.GetEntry(prefix)
		if !exists {
			return fmt.Errorf("key %s is not accepted by the configuration", prefix)
		}
		out[prefix] = fmt.Sprintf("%v", val)
		return nil
	}
	return nil
}

func readJSON(config *puzzle.Config, raw []byte) (map[string]string, error) {
	out := make(map[string]string)
	var data interface{}
	decoder := json.NewDecoder(bytes.NewReader(raw))
	decoder.UseNumber() // keep number precision
	if err := decoder.Decode(&data); err != nil {
		return out, err
	}
	return out, recursiveJSON(config, data, "", out)
}

func ReadJSONRaw(config *puzzle.Config, raw []byte) error {
	out, err := readJSON(config, raw)
	if err != nil {
		return err
	}
	// inject in config
	for k, v := range out {
		// check if the key is accepted by the config
		if entry, exists := config.GetEntry(k); exists {
			// set the value in the config
			if err := entry.Set(v); err != nil {
				return fmt.Errorf("error while setting key %s with value %v (%T): %v",
					k, v, v, err)
			}
		}
	}
	return nil
}

func realSingleEntry(config *puzzle.Config, files []string) error {
	failures := make(map[string]error)
	for _, file := range files {
		raw, err := os.ReadFile(file)
		if err != nil {
			// we ignore and go to the next file
			failures[file] = err
			continue
		}
		if err := ReadJSONRaw(config, raw); err != nil {
			return err
		}
		return nil
	}

	if len(failures) == 0 {
		return fmt.Errorf("no file read")
	}

	msg := "Fail to parse json files:\n"
	for k, v := range failures {
		msg += fmt.Sprintf("  - %s: %v\n", k, v)
	}
	return fmt.Errorf("%s", msg)
}

func ReadJSON(config *puzzle.Config) error {
	for entry := range config.Entries() {
		// filter on config files variables
		if entry.GetMetadata().IsConfigFile {
			// try to cast the entry (this is controlled by the library so it should work)
			files, err := puzzle.Get[[]string](config, entry.GetKey())
			if err != nil {
				return fmt.Errorf("error while retrieving config file variable %s: %v", entry.GetKey(), err)
			}
			if err := realSingleEntry(config, files); err != nil {
				return err
			}
		}
	}

	return nil
}

func ToJSON(config *puzzle.Config) ([]byte, error) {
	tmp := make(map[string]interface{})
	for entry := range config.Entries() {
		k := entry.GetKey()
		tmp0 := tmp
		if strings.Contains(k, config.NestingSeparator) {
			keys := strings.Split(k, config.NestingSeparator)
			// tmp0 := tmp
			for i := 0; i < len(keys)-1; i++ {
				if _, ok := tmp0[keys[i]]; !ok {
					tmp0[keys[i]] = make(map[string]interface{})
				}
				tmp0 = tmp0[keys[i]].(map[string]interface{})
			}
			k = keys[len(keys)-1]
		}

		// in the json case we must return the real typed value in general
		// but for some specific cases, we need to use the string representation
		switch v := entry.GetValue().(type) {
		case time.Duration:
			tmp0[k] = entry.String()
		case []byte:
			tmp0[k] = entry.String()
		default:
			tmp0[k] = v
		}
	}
	return json.Marshal(tmp)
}
