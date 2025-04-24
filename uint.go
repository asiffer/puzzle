package puzzle

import "strconv"

// Size of int and uint (32 or 64 supported)
const uintBits = 32 << (^uint(0) >> 63) // 32 on 32-bit, 64 on 64-bit

var UintConverter = newConverter(uintFromString)

func uintFromString(entry *Entry[uint], stringValue string) error {
	value, err := strconv.ParseUint(stringValue, 10, uintBits)
	if err != nil {
		return err
	}
	*entry.ValueP = uint(value)
	entry.Value = uint(value)
	return nil
}
