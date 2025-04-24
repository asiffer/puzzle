package puzzle

import (
	"strings"
)

var StringSliceConverter = newConverter(stringSliceFromString)

func stringSliceFromString(entry *Entry[[]string], stringValue string) error {
	value := []string{}
	// strings.Split returns a slice of length 1 if the separator is not found
	// where the unique element is the original string
	// however, if the element is empty we should rather have a slice of size 0
	if stringValue != "" {
		value = strings.Split(
			strings.TrimSpace(stringValue),
			entry.Metadata.SliceSeparator,
		)
	}

	*entry.ValueP = value
	entry.Value = value
	return nil
}
