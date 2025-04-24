package puzzle

import (
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

func testDurationConverter(t *testing.T, value uint64) {
	d := time.Duration(value)
	if err := testConverter(DurationConverter, d.String(), d); err != nil {
		t.Error(err)
	}
}

func FuzzDurationConverter(f *testing.F) {
	for i := 0; i < FUZZ_SIZE; i++ {
		f.Add(gofakeit.Uint64())
	}

	f.Fuzz(testDurationConverter)
}
