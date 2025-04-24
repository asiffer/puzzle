package puzzle

import "strconv"

var Int8Converter = newConverter(int8FromString)

func int8FromString(entry *Entry[int8], stringValue string) error {
	value, err := strconv.ParseInt(stringValue, 10, 8)
	if err != nil {
		return err
	}
	*entry.ValueP = int8(value)
	entry.Value = int8(value)
	return nil
}
