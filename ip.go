package puzzle

import (
	"fmt"
	"net"
)

var IPConverter = newConverter(ipFromString)

func ipFromString(entry *Entry[net.IP], stringValue string) error {
	value := net.ParseIP(stringValue)
	if value == nil {
		return &InvalidValueError{Key: entry.Key, Value: stringValue, Err: fmt.Errorf("invalid IP address")}
	}
	*entry.ValueP = value
	entry.Value = value
	return nil
}
