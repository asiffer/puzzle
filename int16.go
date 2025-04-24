package puzzle

import "strconv"

var Int16Converter = newConverter(int16FromString)

func int16FromString(entry *Entry[int16], stringValue string) error {
	value, err := strconv.ParseInt(stringValue, 10, 16)
	if err != nil {
		return err
	}
	*entry.ValueP = int16(value)
	entry.Value = int16(value)
	return nil
}
