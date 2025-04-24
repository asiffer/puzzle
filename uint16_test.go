package puzzle

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
)

func testUint16Converter(t *testing.T, value uint16) {
	// for _, value := range testValues {
	if err := testConverter(Uint16Converter, fmt.Sprintf("%d", value), value); err != nil {
		t.Error(err)
	}
	// }
}

func FuzzUint16Converter(f *testing.F) {
	for i := 0; i < FUZZ_SIZE; i++ {
		f.Add(gofakeit.Uint16())
	}

	f.Fuzz(testUint16Converter)
}
