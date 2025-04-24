package puzzle

import (
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

var BytesConverter = newConverter(bytesFromString)

func bytesFromString(entry *Entry[[]byte], stringValue string) error {
	var value []byte
	var err error

	switch entry.Metadata.Format {
	case "base64":
		value, err = base64.StdEncoding.DecodeString(stringValue)
		if err != nil {
			return fmt.Errorf("error while decoding as base64: %v", err)
		}
	case "base32":
		value, err = base32.StdEncoding.DecodeString(stringValue)
		if err != nil {
			return fmt.Errorf("error while decoding as base32: %v", err)
		}
	default: // hex
		value, err = hex.DecodeString(stringValue)
		if err != nil {
			return fmt.Errorf("error while decoding as hex: %v", err)
		}
	}

	*entry.ValueP = value
	entry.Value = value
	return nil
}
