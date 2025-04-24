package puzzle

import "strconv"

var Int32Converter = newConverter(int32FromString)

func int32FromString(entry *Entry[int32], stringValue string) error {
	value, err := strconv.ParseInt(stringValue, 10, 32)
	if err != nil {
		return err
	}
	*entry.ValueP = int32(value)
	entry.Value = int32(value)
	return nil
}
