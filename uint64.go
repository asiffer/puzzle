package puzzle

import "strconv"

var Uint64Converter = newConverter(uint64FromString)

func uint64FromString(entry *Entry[uint64], stringValue string) error {
	value, err := strconv.ParseUint(stringValue, 10, 64)
	if err != nil {
		return err
	}
	*entry.ValueP = value
	entry.Value = value
	return nil
}
