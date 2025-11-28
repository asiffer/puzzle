package flagset

import (
	"flag"
	"testing"

	"github.com/asiffer/puzzle/frontendtesting"
	"github.com/brianvoe/gofakeit/v7"
)

func testBuild(t *testing.T, i int) {
	gofakeit.Seed(i)
	config, initial := frontendtesting.RandomConfig()
	config2, values := frontendtesting.RandomConfig()
	// toFlags(config2)

	fs, err := Build(config, "test", flag.PanicOnError)
	if err != nil {
		t.Fatalf("error building flagset: %v", err)
	}
	if err := fs.Parse(config2.ToFlags(false)); err != nil {
		t.Fatalf("error parsing flags: %v", err)
	}

	if err := initial.Compare(values); err != nil {
		t.Error(err)
	}
}

func testBuildNested(t *testing.T, i int) {
	gofakeit.Seed(i)
	config, initial := frontendtesting.RandomNestedConfig()
	config2, values := frontendtesting.RandomNestedConfig()

	fs := flag.NewFlagSet("test", flag.PanicOnError)
	if err := Populate(config, fs); err != nil {
		t.Fatalf("error building flagset: %v", err)
	}
	if err := fs.Parse(config2.ToFlags(false)); err != nil {
		t.Fatalf("error parsing flags: %v", err)
	}

	if err := initial.Compare(values); err != nil {
		t.Error(err)
	}
}

// func testBuildShort(t *testing.T, i int) {
// 	gofakeit.Seed(i)
// 	config, initial := frontendtesting.RandomConfigWithShort()
// 	config2, values := frontendtesting.RandomConfigWithShort()
// 	// toFlags(config2)

// 	fs, err := Build(config, "test", flag.PanicOnError)
// 	if err != nil {
// 		t.Fatalf("error building flagset: %v", err)
// 	}
// 	flags := config2.ToFlags(true)
// 	t.Logf("FLAGS: %v\n", flags)
// 	if err := fs.Parse(config2.ToFlags(true)); err != nil {
// 		t.Fatalf("error parsing flags: %v", err)
// 	}

// 	if err := initial.Compare(values); err != nil {
// 		t.Error(err)
// 	}
// }

func FuzzBuild(f *testing.F) {
	for i := range 200 {
		f.Add(i)
	}
	f.Fuzz(testBuild)
}

func FuzzBuildNested(f *testing.F) {
	for i := range 200 {
		f.Add(i)
	}
	f.Fuzz(testBuildNested)
}

// func FuzzBuildShort(f *testing.F) {
// 	for i := range 200 {
// 		f.Add(i)
// 	}
// 	f.Fuzz(testBuildShort)
// }
