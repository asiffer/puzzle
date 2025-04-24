package puzzle

import "strconv"

var Int64Converter = newConverter(int64FromString)

func int64FromString(entry *Entry[int64], stringValue string) error {
	value, err := strconv.ParseInt(stringValue, 10, 64)
	if err != nil {
		return err
	}
	*entry.ValueP = value
	entry.Value = value
	return nil
}
