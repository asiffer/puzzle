package puzzle

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
)

func testFloat32Converter(t *testing.T, value float32) {
	if err := testConverter(Float32Converter, fmt.Sprintf("%v", value), value); err != nil {
		t.Error(err)
	}

}

func FuzzFloat32Converter(f *testing.F) {
	for i := 0; i < FUZZ_SIZE; i++ {
		f.Add(gofakeit.Float32())
	}

	f.Fuzz(testFloat32Converter)
}
