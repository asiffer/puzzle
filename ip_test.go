package puzzle

import (
	"fmt"
	"net"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
)

func testIP4Converter(t *testing.T, a, b, c, d byte) {
	ip := net.IPv4(a, b, c, d)
	if err := testConverterAny(IPConverter, fmt.Sprintf("%v", ip), ip, ip.Equal); err != nil {
		t.Error(err)
	}
}

func testIP6Converter(t *testing.T, ipraw []byte) {
	ip6 := net.IP(ipraw).To16()
	if err := testConverterAny(IPConverter, fmt.Sprintf("%v", ip6), ip6, ip6.Equal); err != nil {
		t.Error(err)
	}
}

func FuzzIP4Converter(f *testing.F) {
	for i := 0; i < FUZZ_SIZE; i++ {
		f.Add(
			gofakeit.Uint8(),
			gofakeit.Uint8(),
			gofakeit.Uint8(),
			gofakeit.Uint8())
	}

	f.Fuzz(testIP4Converter)
}

func FuzzIP6Converter(f *testing.F) {
	for i := 0; i < FUZZ_SIZE; i++ {
		ipraw, err := randomBytes(16)
		if err != nil {
			f.Fatal(err)
		}
		f.Add(ipraw)
	}

	f.Fuzz(testIP6Converter)
}
