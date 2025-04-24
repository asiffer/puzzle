package puzzle

import (
	"os"
)

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
		if err := entry.Convert("string", strValue); err != nil {
			return err
		}
	}
	return nil
}
