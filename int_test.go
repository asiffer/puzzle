package puzzle

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
)

func testIntConverter(t *testing.T, value int) {
	// for _, value := range testValues {
	if err := testConverter(IntConverter, fmt.Sprintf("%d", value), value); err != nil {
		t.Error(err)
	}
	// }
}

func FuzzIntConverter(f *testing.F) {
	for i := 0; i < FUZZ_SIZE; i++ {
		f.Add(gofakeit.Int())
	}

	f.Fuzz(testIntConverter)
}
