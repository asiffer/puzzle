package puzzle

import "strconv"

var BoolConverter = newConverter(boolFromString)

func boolFromString(entry *Entry[bool], stringValue string) error {
	value, err := strconv.ParseBool(stringValue)
	if err != nil {
		return err
	}
	*entry.ValueP = value
	entry.Value = value
	return nil
}
