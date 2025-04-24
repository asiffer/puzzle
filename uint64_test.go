package puzzle

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
)

func testUint64Converter(t *testing.T, value uint64) {
	// for _, value := range testValues {
	if err := testConverter(Uint64Converter, fmt.Sprintf("%d", value), value); err != nil {
		t.Error(err)
	}
	// }
}

func FuzzUint64Converter(f *testing.F) {
	for i := 0; i < FUZZ_SIZE; i++ {
		f.Add(gofakeit.Uint64())
	}

	f.Fuzz(testUint64Converter)
}
