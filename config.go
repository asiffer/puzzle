package puzzle

import "fmt"

type ConfigFilter func(key string) bool

const DEFAULT_NESTING_SEPARATOR = "."

// Config is a structure that plays the role of a namespace if you need to manage
// several configurations
type Config struct {
	NestingSeparator string
	entries          map[string]EntryInterface
	filters          []ConfigFilter
}

// NewConfig inits a new config object
func NewConfig() *Config {
	return &Config{
		NestingSeparator: DEFAULT_NESTING_SEPARATOR,
		entries:          make(map[string]EntryInterface),
		filters:          make([]ConfigFilter, 0),
	}
}

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

	go func() {
		defer close(ch)

		for _, entry := range config.entries {
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
					// falbback to long flag
					out = append(out, fmt.Sprintf("--%s=%s", fn, e.String()))
				}
			default:
				out = append(out, "-"+sfn, e.String())
			}
		} else {
			out = append(out, fmt.Sprintf("--%s=%s", fn, e.String()))
		}
	}
	return out
}
