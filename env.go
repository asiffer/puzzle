package puzzle

import (
	"os"
)

// ReadEnv reads environment variables and updates the Config entries accordingly.
// It looks for each entry's EnvName metadata and tries to find the corresponding
// environment variable. If found, it converts the value using the string converter
// (Set method).
func ReadEnv(c *Config) error {
	for entry := range c.Entries() {
		name := entry.GetMetadata().EnvName
		if name == "" {
			continue
		}
		strValue, exists := os.LookupEnv(name)
		if !exists {
			continue
		}
		if err := entry.Set(strValue); err != nil {
			return err
		}
	}
	return nil
}
