package puzzle

import "strconv"

var Float64Converter = newConverter(float64FromString)

func float64FromString(entry *Entry[float64], stringValue string) error {
	value, err := strconv.ParseFloat(stringValue, 64)
	if err != nil {
		return err
	}
	*entry.ValueP = value
	entry.Value = value
	return nil
}
