package urfave3

import (
	"context"
	"testing"

	"github.com/asiffer/puzzle"
	"github.com/asiffer/puzzle/frontendtesting"
	"github.com/brianvoe/gofakeit/v7"

	"github.com/urfave/cli/v3"
)

func testBuild(
	config *puzzle.Config,
	config2 *puzzle.Config,
	values *frontendtesting.AllTypes,
	values2 *frontendtesting.AllTypes,
	useShort bool,
) error {
	flags, err := Build(config)
	if err != nil {
		return err
	}
	cmd := &cli.Command{
		Name:                   "test",
		Flags:                  flags,
		UseShortOptionHandling: true,
		Action: func(ctx context.Context, c *cli.Command) error {
			return nil
		},
	}

	args := append([]string{"test"}, config2.ToFlags(useShort)...)
	if err := cmd.Run(context.Background(), args); err != nil {
		return err
	}

	return values.Compare(values2)
}

func testBuildLong(t *testing.T, i int) {
	gofakeit.Seed(0)
	config, initial := frontendtesting.RandomConfig()
	config2, values := frontendtesting.RandomConfig()

	if err := testBuild(config, config2, initial, values, false); err != nil {
		t.Fatalf("error building flagset: %v", err)
	}
}

func testBuildShort(t *testing.T, i int) {
	gofakeit.Seed(i)
	config, initial := frontendtesting.RandomConfigWithShort()
	// if a boolean flag is true by default, urfave3 does not accept to
	// pass -b in the cli without argument
	for initial.B {
		config, initial = frontendtesting.RandomConfigWithShort()
	}
	config2, values := frontendtesting.RandomConfigWithShort()

	if err := testBuild(config, config2, initial, values, true); err != nil {
		t.Fatalf("error building flagset: %v", err)
	}
}

func FuzzTestBuildLong(f *testing.F) {
	for i := range 200 {
		f.Add(i)
	}
	f.Fuzz(testBuildLong)
}

func FuzzTestBuildShort(f *testing.F) {
	for i := range 200 {
		f.Add(i)
	}
	f.Fuzz(testBuildShort)
}

func TestEntryCreator(t *testing.T) {
	ec := EntryCreator[bool]{}
	if ec.ToString(true) != "true" {
		t.Errorf("expected 'true', got '%s'", ec.ToString(true))
	}
}
