package urfave3

import (
	"context"
	"testing"

	"github.com/asiffer/puzzle"
	"github.com/asiffer/puzzle/frontendtesting"
	"github.com/brianvoe/gofakeit/v7"

	"github.com/urfave/cli/v3"
)

func randomConfig() (*puzzle.Config, *frontendtesting.AllTypes) {
	config := puzzle.NewConfig()
	defaultValues := frontendtesting.RandomValues()

	puzzle.DefineVar(config, "bool", &defaultValues.B)
	puzzle.DefineVar(config, "string", &defaultValues.S)
	puzzle.DefineVar(config, "int", &defaultValues.I)
	puzzle.DefineVar(config, "int8", &defaultValues.I8)
	puzzle.DefineVar(config, "int16", &defaultValues.I16)
	puzzle.DefineVar(config, "int32", &defaultValues.I32)
	puzzle.DefineVar(config, "int64", &defaultValues.I64)
	puzzle.DefineVar(config, "uint", &defaultValues.U)
	puzzle.DefineVar(config, "uint8", &defaultValues.U8)
	puzzle.DefineVar(config, "uint16", &defaultValues.U16)
	puzzle.DefineVar(config, "uint32", &defaultValues.U32)
	puzzle.DefineVar(config, "uint64", &defaultValues.U64)
	puzzle.DefineVar(config, "float32", &defaultValues.F32)
	puzzle.DefineVar(config, "float64", &defaultValues.F64)
	puzzle.DefineVar(config, "duration", &defaultValues.D)
	puzzle.DefineVar(config, "ip", &defaultValues.IP)
	puzzle.DefineVar(config, "bytes", &defaultValues.Bytes, puzzle.WithFormat("base64"))
	puzzle.DefineVar(config, "string-slice", &defaultValues.SS)

	return config, defaultValues
}

func randomConfigWithShort() (*puzzle.Config, *frontendtesting.AllTypes) {
	config := puzzle.NewConfig()
	defaultValues := frontendtesting.RandomValues()

	puzzle.DefineVar(config, "bool", &defaultValues.B, puzzle.WithShortFlagName("b"))
	puzzle.DefineVar(config, "string", &defaultValues.S, puzzle.WithShortFlagName("c"))
	puzzle.DefineVar(config, "int", &defaultValues.I, puzzle.WithShortFlagName("i"))
	puzzle.DefineVar(config, "int8", &defaultValues.I8, puzzle.WithShortFlagName("j"))
	puzzle.DefineVar(config, "int16", &defaultValues.I16, puzzle.WithShortFlagName("k"))
	puzzle.DefineVar(config, "int32", &defaultValues.I32, puzzle.WithShortFlagName("l"))
	puzzle.DefineVar(config, "int64", &defaultValues.I64, puzzle.WithShortFlagName("m"))
	puzzle.DefineVar(config, "uint", &defaultValues.U, puzzle.WithShortFlagName("n"))
	puzzle.DefineVar(config, "uint8", &defaultValues.U8, puzzle.WithShortFlagName("o"))
	puzzle.DefineVar(config, "uint16", &defaultValues.U16, puzzle.WithShortFlagName("p"))
	puzzle.DefineVar(config, "uint32", &defaultValues.U32, puzzle.WithShortFlagName("q"))
	puzzle.DefineVar(config, "uint64", &defaultValues.U64, puzzle.WithShortFlagName("r"))
	puzzle.DefineVar(config, "float32", &defaultValues.F32, puzzle.WithShortFlagName("s"))
	puzzle.DefineVar(config, "float64", &defaultValues.F64, puzzle.WithShortFlagName("t"))
	puzzle.DefineVar(config, "duration", &defaultValues.D, puzzle.WithShortFlagName("d"))
	puzzle.DefineVar(config, "ip", &defaultValues.IP, puzzle.WithShortFlagName("e"))
	puzzle.DefineVar(config, "bytes", &defaultValues.Bytes, puzzle.WithFormat("base64"), puzzle.WithShortFlagName("f"))
	puzzle.DefineVar(config, "string-slice", &defaultValues.SS, puzzle.WithShortFlagName("g"))

	return config, defaultValues
}

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
	config, initial := randomConfig()
	config2, values := randomConfig()

	if err := testBuild(config, config2, initial, values, false); err != nil {
		t.Fatalf("error building flagset: %v", err)
	}
}

func testBuildShort(t *testing.T, i int) {
	gofakeit.Seed(i)
	config, initial := randomConfigWithShort()
	// if a boolean flag is true by default, urfave3 does not accept to
	// pass -b in the cli without argument
	for initial.B {
		config, initial = randomConfigWithShort()
	}
	config2, values := randomConfigWithShort()

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
