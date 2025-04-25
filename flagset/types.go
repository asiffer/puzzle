package flagset

import (
	"flag"
	"net"
	"time"

	"github.com/asiffer/puzzle"
)

// during the init phase we inject the frontend into every converter
func init() {
	// default
	puzzle.BoolConverter.Register(FlagFrontend, puzzle.ConvertCallbackFactory1(ignoreFlag(boolFlag)))
	puzzle.DurationConverter.Register(FlagFrontend, puzzle.ConvertCallbackFactory1(ignoreFlag(durationFlag)))
	puzzle.Float64Converter.Register(FlagFrontend, puzzle.ConvertCallbackFactory1(ignoreFlag(float64Flag)))
	puzzle.IntConverter.Register(FlagFrontend, puzzle.ConvertCallbackFactory1(ignoreFlag(intFlag)))
	puzzle.Int64Converter.Register(FlagFrontend, puzzle.ConvertCallbackFactory1(ignoreFlag(int64Flag)))
	puzzle.StringConverter.Register(FlagFrontend, puzzle.ConvertCallbackFactory1(ignoreFlag(stringFlag)))
	puzzle.UintConverter.Register(FlagFrontend, puzzle.ConvertCallbackFactory1(ignoreFlag(uintFlag)))
	puzzle.Uint64Converter.Register(FlagFrontend, puzzle.ConvertCallbackFactory1(ignoreFlag(uint64Flag)))
	// extended
	puzzle.BytesConverter.Register(FlagFrontend, puzzle.ConvertCallbackFactory1(ignoreFlag(bytesFlag)))
	puzzle.IPConverter.Register(FlagFrontend, puzzle.ConvertCallbackFactory1(ignoreFlag(ipFlag)))
	puzzle.Int8Converter.Register(FlagFrontend, puzzle.ConvertCallbackFactory1(ignoreFlag(int8Flag)))
	puzzle.Int16Converter.Register(FlagFrontend, puzzle.ConvertCallbackFactory1(ignoreFlag(int16Flag)))
	puzzle.Int32Converter.Register(FlagFrontend, puzzle.ConvertCallbackFactory1(ignoreFlag(int32Flag)))
	puzzle.Uint8Converter.Register(FlagFrontend, puzzle.ConvertCallbackFactory1(ignoreFlag(uint8Flag)))
	puzzle.Uint16Converter.Register(FlagFrontend, puzzle.ConvertCallbackFactory1(ignoreFlag(uint16Flag)))
	puzzle.Uint32Converter.Register(FlagFrontend, puzzle.ConvertCallbackFactory1(ignoreFlag(uint32Flag)))
	puzzle.StringSliceConverter.Register(FlagFrontend, puzzle.ConvertCallbackFactory1(ignoreFlag(stringSliceFlag)))
	puzzle.Float32Converter.Register(FlagFrontend, puzzle.ConvertCallbackFactory1(ignoreFlag(float32Flag)))
}

// supported types

type flagFun[T any] func(entry *puzzle.Entry[T], fs *flag.FlagSet) error

func ignoreFlag[T any](f flagFun[T]) flagFun[T] {
	return func(entry *puzzle.Entry[T], fs *flag.FlagSet) error {
		if entry.Metadata.FlagName == "" {
			return nil
		}
		return f(entry, fs)
	}
}

func boolFlag(entry *puzzle.Entry[bool], fs *flag.FlagSet) error {
	fs.BoolVar(entry.ValueP, entry.Metadata.FlagName, entry.Value, entry.Metadata.Description)
	return nil
}

func durationFlag(entry *puzzle.Entry[time.Duration], fs *flag.FlagSet) error {
	fs.DurationVar(entry.ValueP, entry.Metadata.FlagName, entry.Value, entry.Metadata.Description)
	return nil
}

func float64Flag(entry *puzzle.Entry[float64], fs *flag.FlagSet) error {
	fs.Float64Var(entry.ValueP, entry.Metadata.FlagName, entry.Value, entry.Metadata.Description)
	return nil
}

func intFlag(entry *puzzle.Entry[int], fs *flag.FlagSet) error {
	fs.IntVar(entry.ValueP, entry.Metadata.FlagName, entry.Value, entry.Metadata.Description)
	return nil
}

func int64Flag(entry *puzzle.Entry[int64], fs *flag.FlagSet) error {
	fs.Int64Var(entry.ValueP, entry.Metadata.FlagName, entry.Value, entry.Metadata.Description)
	return nil
}

func stringFlag(entry *puzzle.Entry[string], fs *flag.FlagSet) error {
	fs.StringVar(entry.ValueP, entry.Metadata.FlagName, entry.Value, entry.Metadata.Description)
	return nil
}

func uintFlag(entry *puzzle.Entry[uint], fs *flag.FlagSet) error {
	fs.UintVar(entry.ValueP, entry.Metadata.FlagName, entry.Value, entry.Metadata.Description)
	return nil
}

func uint64Flag(entry *puzzle.Entry[uint64], fs *flag.FlagSet) error {
	fs.Uint64Var(entry.ValueP, entry.Metadata.FlagName, entry.Value, entry.Metadata.Description)
	return nil
}

// extended

func bytesFlag(entry *puzzle.Entry[[]byte], fs *flag.FlagSet) error {
	fs.Var(entry, entry.Metadata.FlagName, entry.Metadata.Description)
	return nil
}

func ipFlag(entry *puzzle.Entry[net.IP], fs *flag.FlagSet) error {
	fs.TextVar(entry.ValueP, entry.Metadata.FlagName, entry.Value, entry.Metadata.Description)
	return nil
}

func int8Flag(entry *puzzle.Entry[int8], fs *flag.FlagSet) error {
	fs.Var(entry, entry.Metadata.FlagName, entry.Metadata.Description)
	return nil
}

func int16Flag(entry *puzzle.Entry[int16], fs *flag.FlagSet) error {
	fs.Var(entry, entry.Metadata.FlagName, entry.Metadata.Description)
	return nil
}

func int32Flag(entry *puzzle.Entry[int32], fs *flag.FlagSet) error {
	fs.Var(entry, entry.Metadata.FlagName, entry.Metadata.Description)
	return nil
}

func uint8Flag(entry *puzzle.Entry[uint8], fs *flag.FlagSet) error {
	fs.Var(entry, entry.Metadata.FlagName, entry.Metadata.Description)
	return nil
}

func uint16Flag(entry *puzzle.Entry[uint16], fs *flag.FlagSet) error {
	fs.Var(entry, entry.Metadata.FlagName, entry.Metadata.Description)
	return nil
}

func uint32Flag(entry *puzzle.Entry[uint32], fs *flag.FlagSet) error {
	fs.Var(entry, entry.Metadata.FlagName, entry.Metadata.Description)
	return nil
}

func stringSliceFlag(entry *puzzle.Entry[[]string], fs *flag.FlagSet) error {
	fs.Var(entry, entry.Metadata.FlagName, entry.Metadata.Description)
	return nil
}

func float32Flag(entry *puzzle.Entry[float32], fs *flag.FlagSet) error {
	fs.Var(entry, entry.Metadata.FlagName, entry.Metadata.Description)
	return nil
}
