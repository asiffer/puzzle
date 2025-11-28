package puzzle

import (
	"net/url"
)

// ReadForm reads the form values from the provided url.Values and updates them in the Config.
// It relies on the string frontend to convert the values (what the Set method does).
func ReadForm(c *Config, form url.Values) error {
	for entry := range c.Entries() {
		name := entry.GetKey()
		if name == "" {
			continue
		}
		if !form.Has(name) {
			continue
		}
		values := form[name]
		if len(values) == 1 {
			if err := entry.Set(values[0]); err != nil {
				return err
			}
		} else if len(values) > 1 {
			if err := entry.Set(join(values, entry)); err != nil {
				return err
			}
		}
	}
	return nil
}
