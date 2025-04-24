package puzzle

import "strconv"

var Float32Converter = newConverter(float32FromString)

func float32FromString(entry *Entry[float32], stringValue string) error {
	value, err := strconv.ParseFloat(stringValue, 32)
	if err != nil {
		return err
	}
	*entry.ValueP = float32(value)
	entry.Value = float32(value)
	return nil
}
