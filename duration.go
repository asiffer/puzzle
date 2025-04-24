package puzzle

import (
	"time"
)

var DurationConverter Converter[time.Duration] = newConverter(durationFromString)

func durationFromString(entry *Entry[time.Duration], stringValue string) error {
	value, err := time.ParseDuration(stringValue)
	if err != nil {
		return err
	}
	*entry.ValueP = value
	entry.Value = value
	return nil
}
