package pflagset

import (
	"net"
	"time"

	"github.com/asiffer/puzzle"
	"github.com/spf13/pflag"
)

func init() {
	puzzle.BoolConverter.Register(PFlagFrontend, puzzle.ConvertCallbackFactory1(boolFlag))
	puzzle.DurationConverter.Register(PFlagFrontend, puzzle.ConvertCallbackFactory1(autoFlag[time.Duration]))
	puzzle.Float64Converter.Register(PFlagFrontend, puzzle.ConvertCallbackFactory1(autoFlag[float64]))
	puzzle.IntConverter.Register(PFlagFrontend, puzzle.ConvertCallbackFactory1(autoFlag[int]))
	puzzle.Int64Converter.Register(PFlagFrontend, puzzle.ConvertCallbackFactory1(autoFlag[int64]))
	puzzle.StringConverter.Register(PFlagFrontend, puzzle.ConvertCallbackFactory1(autoFlag[string]))
	puzzle.UintConverter.Register(PFlagFrontend, puzzle.ConvertCallbackFactory1(autoFlag[uint]))
	puzzle.Uint64Converter.Register(PFlagFrontend, puzzle.ConvertCallbackFactory1(autoFlag[uint64]))
	puzzle.BytesConverter.Register(PFlagFrontend, puzzle.ConvertCallbackFactory1(autoFlag[[]byte]))
	puzzle.IPConverter.Register(PFlagFrontend, puzzle.ConvertCallbackFactory1(autoFlag[net.IP]))
	puzzle.Int8Converter.Register(PFlagFrontend, puzzle.ConvertCallbackFactory1(autoFlag[int8]))
	puzzle.Int16Converter.Register(PFlagFrontend, puzzle.ConvertCallbackFactory1(autoFlag[int16]))
	puzzle.Int32Converter.Register(PFlagFrontend, puzzle.ConvertCallbackFactory1(autoFlag[int32]))
	puzzle.Uint8Converter.Register(PFlagFrontend, puzzle.ConvertCallbackFactory1(autoFlag[uint8]))
	puzzle.Uint16Converter.Register(PFlagFrontend, puzzle.ConvertCallbackFactory1(autoFlag[uint16]))
	puzzle.Uint32Converter.Register(PFlagFrontend, puzzle.ConvertCallbackFactory1(autoFlag[uint32]))
	puzzle.StringSliceConverter.Register(PFlagFrontend, puzzle.ConvertCallbackFactory1(autoFlag[[]string]))
	puzzle.Float32Converter.Register(PFlagFrontend, puzzle.ConvertCallbackFactory1(autoFlag[float32]))
}

// supported types

func autoFlag[T any](entry *puzzle.Entry[T], fs *pflag.FlagSet) error {
	if entry.Metadata.FlagName != "" {
		if entry.Metadata.ShortFlagName == "" {
			fs.Var(entry, entry.Metadata.FlagName, entry.Metadata.Description)
		} else {
			fs.VarP(entry, entry.Metadata.FlagName, entry.Metadata.ShortFlagName, entry.Metadata.Description)
		}
	}
	return nil
}

// custom flag to support shorthand flags
func boolFlag(entry *puzzle.Entry[bool], fs *pflag.FlagSet) error {
	if entry.Metadata.FlagName != "" {
		if entry.Metadata.ShortFlagName == "" {
			fs.BoolVar(entry.ValueP, entry.Metadata.FlagName, entry.Value, entry.Metadata.Description)
		} else {
			fs.BoolVarP(entry.ValueP, entry.Metadata.FlagName, entry.Metadata.ShortFlagName, entry.Value, entry.Metadata.Description)
		}
	}
	return nil
}
