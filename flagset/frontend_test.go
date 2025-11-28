package flagset

import (
	"flag"
	"testing"

	"github.com/asiffer/puzzle/frontendtesting"
	"github.com/brianvoe/gofakeit/v7"
)

// func randomConfig() (*puzzle.Config, *frontendtesting.AllTypes) {
// 	config := puzzle.NewConfig()
// 	defaultValues := frontendtesting.RandomValues()

// 	puzzle.DefineVar(config, "bool", &defaultValues.B)
// 	puzzle.DefineVar(config, "string", &defaultValues.S)
// 	puzzle.DefineVar(config, "int", &defaultValues.I)
// 	puzzle.DefineVar(config, "int8", &defaultValues.I8)
// 	puzzle.DefineVar(config, "int16", &defaultValues.I16)
// 	puzzle.DefineVar(config, "int32", &defaultValues.I32)
// 	puzzle.DefineVar(config, "int64", &defaultValues.I64)
// 	puzzle.DefineVar(config, "uint", &defaultValues.U)
// 	puzzle.DefineVar(config, "uint8", &defaultValues.U8)
// 	puzzle.DefineVar(config, "uint16", &defaultValues.U16)
// 	puzzle.DefineVar(config, "uint32", &defaultValues.U32)
// 	puzzle.DefineVar(config, "uint64", &defaultValues.U64)
// 	puzzle.DefineVar(config, "float32", &defaultValues.F32)
// 	puzzle.DefineVar(config, "float64", &defaultValues.F64)
// 	puzzle.DefineVar(config, "duration", &defaultValues.D)
// 	puzzle.DefineVar(config, "ip", &defaultValues.IP)
// 	puzzle.DefineVar(config, "bytes", &defaultValues.Bytes, puzzle.WithFormat("base64"))
// 	puzzle.DefineVar(config, "string-slice", &defaultValues.SS)

// 	return config, defaultValues
// }

// func randomConfigWithShort() (*puzzle.Config, *frontendtesting.AllTypes) {
// 	config := puzzle.NewConfig()
// 	defaultValues := frontendtesting.RandomValues()

// 	puzzle.DefineVar(config, "bool", &defaultValues.B, puzzle.WithShortFlagName("b"))
// 	puzzle.DefineVar(config, "string", &defaultValues.S, puzzle.WithShortFlagName("c"))
// 	puzzle.DefineVar(config, "int", &defaultValues.I, puzzle.WithShortFlagName("i"))
// 	puzzle.DefineVar(config, "int8", &defaultValues.I8, puzzle.WithShortFlagName("j"))
// 	puzzle.DefineVar(config, "int16", &defaultValues.I16, puzzle.WithShortFlagName("k"))
// 	puzzle.DefineVar(config, "int32", &defaultValues.I32, puzzle.WithShortFlagName("l"))
// 	puzzle.DefineVar(config, "int64", &defaultValues.I64, puzzle.WithShortFlagName("m"))
// 	puzzle.DefineVar(config, "uint", &defaultValues.U, puzzle.WithShortFlagName("n"))
// 	puzzle.DefineVar(config, "uint8", &defaultValues.U8, puzzle.WithShortFlagName("o"))
// 	puzzle.DefineVar(config, "uint16", &defaultValues.U16, puzzle.WithShortFlagName("p"))
// 	puzzle.DefineVar(config, "uint32", &defaultValues.U32, puzzle.WithShortFlagName("q"))
// 	puzzle.DefineVar(config, "uint64", &defaultValues.U64, puzzle.WithShortFlagName("r"))
// 	puzzle.DefineVar(config, "float32", &defaultValues.F32, puzzle.WithShortFlagName("s"))
// 	puzzle.DefineVar(config, "float64", &defaultValues.F64, puzzle.WithShortFlagName("t"))
// 	puzzle.DefineVar(config, "duration", &defaultValues.D, puzzle.WithShortFlagName("d"))
// 	puzzle.DefineVar(config, "ip", &defaultValues.IP, puzzle.WithShortFlagName("e"))
// 	puzzle.DefineVar(config, "bytes", &defaultValues.Bytes, puzzle.WithFormat("base64"), puzzle.WithShortFlagName("f"))
// 	puzzle.DefineVar(config, "string-slice", &defaultValues.SS, puzzle.WithShortFlagName("g"))

// 	return config, defaultValues
// }

// func nestedKey(key string, prefix string) string {
// 	if len(key) == 0 {
// 		return key
// 	}
// 	return fmt.Sprintf("%s%s%s", prefix, puzzle.DEFAULT_NESTING_SEPARATOR, key)
// }

// func randomNestedConfig() (*puzzle.Config, *frontendtesting.AllTypes) {
// 	config := puzzle.NewConfig()
// 	defaultValues := frontendtesting.RandomValues()

// 	puzzle.DefineVar(config, "ignored-bool", &defaultValues.IB, puzzle.WithoutFlagName())
// 	puzzle.DefineVar(config, "bool", &defaultValues.B)
// 	puzzle.DefineVar(config, "string", &defaultValues.S)
// 	puzzle.DefineVar(config, nestedKey("int", "a"), &defaultValues.I)
// 	puzzle.DefineVar(config, nestedKey("int8", "a"), &defaultValues.I8)
// 	puzzle.DefineVar(config, nestedKey("int16", "a"), &defaultValues.I16)
// 	puzzle.DefineVar(config, nestedKey("int32", "b"), &defaultValues.I32)
// 	puzzle.DefineVar(config, nestedKey("int64", "b"), &defaultValues.I64)
// 	puzzle.DefineVar(config, nestedKey(nestedKey("uint", "c"), "a"), &defaultValues.U)
// 	puzzle.DefineVar(config, nestedKey(nestedKey("uint8", "c"), "a"), &defaultValues.U8)
// 	puzzle.DefineVar(config, nestedKey(nestedKey("uint16", "d"), "a"), &defaultValues.U16)
// 	puzzle.DefineVar(config, nestedKey(nestedKey("uint32", "c"), "b"), &defaultValues.U32)
// 	puzzle.DefineVar(config, nestedKey(nestedKey("uint64", "x"), "y"), &defaultValues.U64)
// 	puzzle.DefineVar(config, "float32", &defaultValues.F32)
// 	puzzle.DefineVar(config, "float64", &defaultValues.F64)
// 	puzzle.DefineVar(config, "duration", &defaultValues.D)
// 	puzzle.DefineVar(config, "ip", &defaultValues.IP)
// 	puzzle.DefineVar(config, "bytes", &defaultValues.Bytes, puzzle.WithFormat("base64"))
// 	puzzle.DefineVar(config, "string-slice", &defaultValues.SS)

// 	return config, defaultValues
// }

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
