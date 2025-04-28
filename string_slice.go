package puzzle

import (
	"strings"
)

var StringSliceConverter = newConverter(stringSliceFromString)

func stringSliceFromString(entry *Entry[[]string], stringValue string) error {
	value := []string{}
	var err error

	if stringValue != "" {
		reader := entry.csvReader(strings.NewReader(stringValue))
		value, err = reader.Read()
		if err != nil {
			return err
		}
	}

	*entry.ValueP = value
	entry.Value = value
	return nil
}
