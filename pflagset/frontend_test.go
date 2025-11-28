package pflagset

import (
	"testing"

	"github.com/asiffer/puzzle"
	"github.com/asiffer/puzzle/frontendtesting"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/spf13/pflag"
)

func baseTestBuild(
	config *puzzle.Config,
	config2 *puzzle.Config,
	values *frontendtesting.AllTypes,
	values2 *frontendtesting.AllTypes,
	useShort bool,
) error {
	fs, err := Build(config, "test", pflag.PanicOnError)
	if err != nil {
		return err
	}
	if err := fs.Parse(config2.ToFlags(useShort)); err != nil {
		return err
	}

	return values.Compare(values2)
}

func testBuild(t *testing.T, i int) {
	gofakeit.Seed(i)
	config, initial := frontendtesting.RandomConfig()
	config2, values := frontendtesting.RandomConfig()

	if err := baseTestBuild(config, config2, initial, values, false); err != nil {
		t.Fatalf("error building flagset: %v", err)
	}
}

func testBuildShort(t *testing.T, i int) {
	gofakeit.Seed(i)
	config, initial := frontendtesting.RandomConfigWithShort()
	config2, values := frontendtesting.RandomConfigWithShort()

	if err := baseTestBuild(config, config2, initial, values, true); err != nil {
		t.Fatalf("error building flagset: %v", err)
	}
}

func FuzzBuild(f *testing.F) {
	for i := range 200 {
		f.Add(i)
	}
	f.Fuzz(testBuild)
}

func FuzzBuildShort(f *testing.F) {
	for i := range 200 {
		f.Add(i)
	}
	f.Fuzz(testBuildShort)
}
