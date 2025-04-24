package puzzle

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
)

func testInt16Converter(t *testing.T, value int16) {
	// for _, value := range testValues {
	if err := testConverter(Int16Converter, fmt.Sprintf("%d", value), value); err != nil {
		t.Error(err)
	}
	// }
}

func FuzzInt16Converter(f *testing.F) {
	for i := 0; i < FUZZ_SIZE; i++ {
		f.Add(gofakeit.Int16())
	}

	f.Fuzz(testInt16Converter)
}
