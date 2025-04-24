package puzzle

import "strconv"

// Size of int and uint (32 or 64 supported)
const intBits = 32 << (^uint(0) >> 63) // 32 on 32-bit, 64 on 64-bit

var IntConverter = newConverter(intFromString)

func intFromString(entry *Entry[int], stringValue string) error {
	value, err := strconv.ParseInt(stringValue, 10, intBits)
	if err != nil {
		return err
	}
	*entry.ValueP = int(value)
	entry.Value = int(value)
	return nil
}
