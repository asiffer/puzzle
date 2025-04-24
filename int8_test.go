package puzzle

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
)

func testInt8Converter(t *testing.T, value int8) {
	// for _, value := range testValues {
	if err := testConverter(Int8Converter, fmt.Sprintf("%d", value), value); err != nil {
		t.Error(err)
	}
	// }
}

func FuzzInt8Converter(f *testing.F) {
	for i := 0; i < FUZZ_SIZE; i++ {
		f.Add(gofakeit.Int8())
	}

	f.Fuzz(testInt8Converter)
}
