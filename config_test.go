package puzzle

import (
	"testing"

	"github.com/brianvoe/gofakeit/v7"
)

func randomKeys(size int) []string {
	keys := make([]string, size)
	for i := 0; i < size; i++ {
		keys[i] = gofakeit.LetterN(12)
	}
	return keys
}

func TestConfigOnly(t *testing.T) {
	k := 5
	n := k + int(gofakeit.Uint16())

	config := NewConfig()
	keys := randomKeys(int(n))
	for _, k := range keys {
		Define(config, k, gofakeit.Int())
	}

	only := config.Only(keys[:k]...)
	var count int = 0
	for range only.Entries() {
		count++
	}
	if count != k {
		t.Errorf("expected %d entries, got %d", k, count)
	}

}

func TestIgnoringOnly(t *testing.T) {
	k := 5
	n := k + int(gofakeit.Uint16())

	config := NewConfig()
	keys := randomKeys(int(n))
	for _, k := range keys {
		Define(config, k, gofakeit.Int())
	}

	only := config.Ignoring(keys[:k]...)
	var count int = 0
	for range only.Entries() {
		count++
	}
	if count != n-k {
		t.Errorf("expected %d entries, got %d", k, count)
	}

}

func testToFlags(t *testing.T, i int) {
	gofakeit.Seed(i)
	conf, _ := randomConfig()
	flags := conf.ToFlags(false)
	if len(flags) != len(conf.entries) {
		t.Errorf("expected %d flags, got %d", len(conf.entries), len(flags))
	}

	flags = conf.ToFlags(true)
	if len(flags) != len(conf.entries) && len(flags) != (len(conf.entries)-1) {
		t.Errorf("expected %d flags, got %d", len(conf.entries), len(flags))
	}
}

func FuzzToFlags(f *testing.F) {
	for i := range FUZZ_SIZE {
		f.Add(i)
	}
	f.Fuzz(testToFlags)
}

func TestGetEntry(t *testing.T) {
	conf, _ := randomConfig()
	for _, e := range conf.entries {
		if _, ok := conf.GetEntry(e.GetKey()); !ok {
			t.Errorf("expected to find entry %s", e.GetKey())
		}
	}
	if _, ok := conf.GetEntry("nonexistent"); ok {
		t.Errorf("expected not to find entry %s", "nonexistent")
	}
}
