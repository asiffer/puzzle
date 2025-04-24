package puzzle

import (
	"strings"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
)

func fakeStringSlice(v uint8) []string {
	out := make([]string, 0)
	// out = append(out, gofakeit.Name())
	generators := []func() string{
		gofakeit.Adjective,
		gofakeit.Noun,
		gofakeit.Verb,
		gofakeit.Color,
		gofakeit.Word,
		gofakeit.Interjection,
		gofakeit.FarmAnimal,
	}
	for i, g := range generators {
		if (v & (1 << i)) != 0 {
			out = append(out, g())
		}
	}
	return out
}

func testStringSliceConverter(t *testing.T, x uint8) {
	// for _, value := range testValues {
	value := fakeStringSlice(x)
	if err := testConverterAny(StringSliceConverter, strings.Join(value, ","), value, sliceEqualFactory(value)); err != nil {
		t.Error(err)
	}
	// }
}

func FuzzStringSliceConverter(f *testing.F) {
	for i := 0; i < FUZZ_SIZE; i++ {
		f.Add(gofakeit.Uint8())
	}

	f.Fuzz(testStringSliceConverter)
}
