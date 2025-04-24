package puzzle

import "strconv"

var Uint16Converter = newConverter(uint16FromString)

func uint16FromString(entry *Entry[uint16], stringValue string) error {
	value, err := strconv.ParseUint(stringValue, 10, 16)
	if err != nil {
		return err
	}
	*entry.ValueP = uint16(value)
	entry.Value = uint16(value)
	return nil
}
