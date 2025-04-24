package puzzle

import (
	"crypto/rand"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"testing"
)

func randomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func testBytesConverterBase64(t *testing.T, value []byte) {
	mod := func(e *Entry[[]byte]) {
		e.Metadata.Format = "base64"
	}
	if err := testConverterAny(BytesConverter, base64.StdEncoding.EncodeToString(value), value, sliceEqualFactory(value), mod); err != nil {
		t.Error(err)
	}

}

func testBytesConverterBase32(t *testing.T, value []byte) {
	mod := func(e *Entry[[]byte]) {
		e.Metadata.Format = "base32"
	}
	if err := testConverterAny(BytesConverter, base32.StdEncoding.EncodeToString(value), value, sliceEqualFactory(value), mod); err != nil {
		t.Error(err)
	}

}

func testBytesConverterHex(t *testing.T, value []byte) {
	mod := func(e *Entry[[]byte]) {
		e.Metadata.Format = "hex"
	}
	if err := testConverterAny(BytesConverter, hex.EncodeToString(value), value, sliceEqualFactory(value), mod); err != nil {
		t.Error(err)
	}

}

func FuzzBytesConverterBase64(f *testing.F) {
	for range FUZZ_SIZE {
		b, err := randomBytes(256)
		if err != nil {
			f.Fatal(err)
		}
		f.Add(b)
	}

	f.Fuzz(testBytesConverterBase64)
}

func FuzzBytesConverterBase32(f *testing.F) {
	for range FUZZ_SIZE {
		b, err := randomBytes(256)
		if err != nil {
			f.Fatal(err)
		}
		f.Add(b)
	}

	f.Fuzz(testBytesConverterBase32)
}

func FuzzBytesConverterHex(f *testing.F) {
	for range FUZZ_SIZE {
		b, err := randomBytes(256)
		if err != nil {
			f.Fatal(err)
		}
		f.Add(b)
	}

	f.Fuzz(testBytesConverterHex)
}
