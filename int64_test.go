package puzzle

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
)

func testInt64Converter(t *testing.T, value int64) {
	// for _, value := range testValues {
	if err := testConverter(Int64Converter, fmt.Sprintf("%d", value), value); err != nil {
		t.Error(err)
	}
	// }
}

func FuzzInt64Converter(f *testing.F) {
	for i := 0; i < FUZZ_SIZE; i++ {
		f.Add(gofakeit.Int64())
	}

	f.Fuzz(testInt64Converter)
}
