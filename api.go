package puzzle

import (
	"fmt"
)

func insert[T any](config *Config, entry *Entry[T], options ...MetadataOption) error {
	if _, exists := config.entries[entry.Key]; exists {
		return fmt.Errorf("key %s already exists", entry.Key)
	}
	// apply options
	for _, option := range options {
		option(entry.Metadata)
	}

	config.entries[entry.Key] = entry
	return nil
}

// Define lets the config object store the Value
func Define[T any](config *Config, key string, defaultValue T, options ...MetadataOption) error {
	entry := newEntry[T](key)
	entry.Value = defaultValue // store the Value locally
	entry.ValueP = &entry.Value
	return insert(config, entry, options...)
}

// DefineVar lets the user store the Value
func DefineVar[T any](config *Config, key string, boundVariable *T, options ...MetadataOption) error {
	entry := newEntry[T](key)
	entry.Value = *boundVariable
	entry.ValueP = boundVariable // keep the provided storage

	return insert(config, entry, options...)
}

// Get retrieves the Value from the config object
func Get[T any](config *Config, key string) (T, error) {
	var out T
	entry, exists := config.entries[key]
	if !exists {
		return out, fmt.Errorf("key %s not found", key)
	}
	Value, ok := entry.GetValue().(T)
	if !ok {
		return out, fmt.Errorf("type mismatch: requested %T, got %T", out, entry.GetValue())
	}
	return Value, nil
}

func DefineConfigFile(config *Config, key string, defaultValue []string, options ...MetadataOption) error {
	// inject a custom option that is not visible by end user
	setConfigFile := func(m *EntryMetadata) {
		m.IsConfigFile = true
	}
	return Define(config, key, defaultValue, append(options, setConfigFile)...)
}
