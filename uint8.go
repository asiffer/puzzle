package puzzle

import "strconv"

var Uint8Converter = newConverter(uint8FromString)

func uint8FromString(entry *Entry[uint8], stringValue string) error {
	value, err := strconv.ParseUint(stringValue, 10, 8)
	if err != nil {
		return err
	}
	*entry.ValueP = uint8(value)
	entry.Value = uint8(value)
	return nil
}
