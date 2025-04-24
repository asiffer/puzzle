package puzzle

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
)

func testUint32Converter(t *testing.T, value uint32) {
	// for _, value := range testValues {
	if err := testConverter(Uint32Converter, fmt.Sprintf("%d", value), value); err != nil {
		t.Error(err)
	}
	// }
}

func FuzzUint32Converter(f *testing.F) {
	for i := 0; i < FUZZ_SIZE; i++ {
		f.Add(gofakeit.Uint32())
	}

	f.Fuzz(testUint32Converter)
}
