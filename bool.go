package puzzle

import "strconv"

var BoolConverter = newConverter(boolFromString)

func boolFromString(entry *Entry[bool], stringValue string) error {
	var value bool
	var err error

	switch stringValue {
	case "on": // we add "on" and "off" for compatibility with HTML forms
		value = true
	case "off":
		value = false
	default:
		value, err = strconv.ParseBool(stringValue)
		if err != nil {
			return err
		}
	}

	*entry.ValueP = value
	entry.Value = value
	return nil
}
