package puzzle

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
)

func testFloat64Converter(t *testing.T, value float64) {
	if err := testConverter(Float64Converter, fmt.Sprintf("%v", value), value); err != nil {
		t.Error(err)
	}

}

func FuzzFloat64Converter(f *testing.F) {
	for i := 0; i < FUZZ_SIZE; i++ {
		f.Add(gofakeit.Float64())
	}

	f.Fuzz(testFloat64Converter)
}
