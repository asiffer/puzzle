package puzzle

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
)

func testInt32Converter(t *testing.T, value int32) {
	// for _, value := range testValues {
	if err := testConverter(Int32Converter, fmt.Sprintf("%d", value), value); err != nil {
		t.Error(err)
	}
	// }
}

func FuzzInt32Converter(f *testing.F) {
	for i := 0; i < FUZZ_SIZE; i++ {
		f.Add(gofakeit.Int32())
	}

	f.Fuzz(testInt32Converter)
}
