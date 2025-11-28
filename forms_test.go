package puzzle

import (
	"net/url"
	"testing"
)

func testForm(t *testing.T, i int) {
	// empty form
	form := url.Values{}
	setter := func(entry EntryInterface, value string) {
		form.Add(entry.GetKey(), value)
	}

	reader := func(c *Config) error {
		return ReadForm(c, form)
	}
	testGenericRead(t, setter, reader, i)
}

func FuzzForm(f *testing.F) {
	for i := 0; i < FUZZ_SIZE; i++ {
		f.Add(i)
	}

	f.Fuzz(testForm)
}
