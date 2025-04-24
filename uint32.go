package puzzle

import "strconv"

var Uint32Converter = newConverter(uint32FromString)

func uint32FromString(entry *Entry[uint32], stringValue string) error {
	value, err := strconv.ParseUint(stringValue, 10, 32)
	if err != nil {
		return err
	}
	*entry.ValueP = uint32(value)
	entry.Value = uint32(value)
	return nil
}
