package puzzle

import (
	"fmt"
	"slices"
)

type ConfigFilter func(key string) bool

const DEFAULT_NESTING_SEPARATOR = "."

// Config is a structure that plays the role of a namespace if you need to manage
// several configurations
type Config struct {
	NestingSeparator string
	entries          map[string]EntryInterface
	filters          []ConfigFilter
	order            []string
}

// NewConfig inits a new config object
func NewConfig() *Config {
	return &Config{
		NestingSeparator: DEFAULT_NESTING_SEPARATOR,
		entries:          make(map[string]EntryInterface),
		filters:          make([]ConfigFilter, 0),
		order:            nil,
	}
}

func (config *Config) orderIfNotProvided() {
	if config.order != nil {
		return
	}

	config.order = make([]string, 0, len(config.entries))
	for key := range config.entries {
		config.order = append(config.order, key)
	}
}

// Sort sorts the entries of the config by their keys in alphabetical order
func (config *Config) Sort() {
	config.order = make([]string, 0, len(config.entries))
	for key := range config.entries {
		config.order = append(config.order, key)
	}
	slices.Sort(config.order)
}

// SortFunc allows to sort the entries of the config using a custom function
// The function should take a slice of strings (the keys) and return a sorted slice of strings
// even if it is sorted in-place.
// This function can be run after a previous sort
func (config *Config) SortFunc(fun func([]string) []string) {
	config.orderIfNotProvided()
	// clone the order to avoid modifying the original
	cpy := make([]string, len(config.order))
	copy(cpy, config.order)
	// assign the new order
	config.order = fun(cpy)
}

// Accept checks if the key is accepted by all filters
func (config *Config) Accept(key string) bool {
	for _, accept := range config.filters {
		if !accept(key) {
			return false
		}
	}
	return true
}

func (config *Config) GetEntry(key string) (EntryInterface, bool) {
	e, ok := config.entries[key]
	return e, ok
}

// Entries iterates over the entries of the config, applying the filters
// This is the single way to access the entrie of the config
func (config *Config) Entries() <-chan EntryInterface {
	ch := make(chan EntryInterface)

	// provide the order if not already set
	config.orderIfNotProvided()

	go func() {
		defer close(ch)

		for _, key := range config.order {
			entry := config.entries[key]
			if !config.Accept(entry.GetKey()) {
				continue
			}
			ch <- entry
		}
	}()

	return ch
}

// Ignoring creates a new view of the same config, ignoring some keys
func (config *Config) Ignoring(keys ...string) *Config {
	return &Config{
		entries: config.entries,
		filters: append(config.filters,
			func(key string) bool {
				for _, k := range keys {
					if key == k {
						return false
					}
				}
				return true
			},
		),
	}
}

// Only creates a new view of the same config, accepting only some keys
func (config *Config) Only(keys ...string) *Config {
	return &Config{
		entries: config.entries,
		filters: append(config.filters,
			func(key string) bool {
				for _, k := range keys {
					if key == k {
						return true
					}
				}
				return false
			},
		),
	}
}

// ToFlags converts the config entries to a slice of command line flags
// If useShort is true, it will use the short flag names if available, otherwise it will use the long flag names
// If the entry is a boolean, it will only add the flag if the value is true
func (config *Config) ToFlags(useShort bool) []string {
	out := make([]string, 0)
	for _, e := range config.entries {
		fn := e.GetMetadata().FlagName
		sfn := e.GetMetadata().ShortFlagName
		if fn == "" {
			continue
		}

		if useShort && sfn != "" {
			switch b := e.GetValue().(type) {
			case bool:
				if b {
					out = append(out, "-"+sfn)
				} else {
					// fallback to long flag
					out = append(out, fmt.Sprintf("--%s=%s", fn, e.String()))
				}
			default:
				out = append(out, "-"+sfn, e.String())
			}
		} else {
			// long case
			out = append(out, fmt.Sprintf("--%s=%s", fn, e.String()))
		}
	}
	return out
}
