package flagset

import (
	"flag"
	"net"
	"time"

	"github.com/asiffer/puzzle"
)

func init() {
	puzzle.BoolConverter.Register(FlagFrontend, puzzle.ConvertCallbackFactory1(autoFlag[bool]))
	puzzle.DurationConverter.Register(FlagFrontend, puzzle.ConvertCallbackFactory1(autoFlag[time.Duration]))
	puzzle.Float64Converter.Register(FlagFrontend, puzzle.ConvertCallbackFactory1(autoFlag[float64]))
	puzzle.IntConverter.Register(FlagFrontend, puzzle.ConvertCallbackFactory1(autoFlag[int]))
	puzzle.Int64Converter.Register(FlagFrontend, puzzle.ConvertCallbackFactory1(autoFlag[int64]))
	puzzle.StringConverter.Register(FlagFrontend, puzzle.ConvertCallbackFactory1(autoFlag[string]))
	puzzle.UintConverter.Register(FlagFrontend, puzzle.ConvertCallbackFactory1(autoFlag[uint]))
	puzzle.Uint64Converter.Register(FlagFrontend, puzzle.ConvertCallbackFactory1(autoFlag[uint64]))
	puzzle.BytesConverter.Register(FlagFrontend, puzzle.ConvertCallbackFactory1(autoFlag[[]byte]))
	puzzle.IPConverter.Register(FlagFrontend, puzzle.ConvertCallbackFactory1(autoFlag[net.IP]))
	puzzle.Int8Converter.Register(FlagFrontend, puzzle.ConvertCallbackFactory1(autoFlag[int8]))
	puzzle.Int16Converter.Register(FlagFrontend, puzzle.ConvertCallbackFactory1(autoFlag[int16]))
	puzzle.Int32Converter.Register(FlagFrontend, puzzle.ConvertCallbackFactory1(autoFlag[int32]))
	puzzle.Uint8Converter.Register(FlagFrontend, puzzle.ConvertCallbackFactory1(autoFlag[uint8]))
	puzzle.Uint16Converter.Register(FlagFrontend, puzzle.ConvertCallbackFactory1(autoFlag[uint16]))
	puzzle.Uint32Converter.Register(FlagFrontend, puzzle.ConvertCallbackFactory1(autoFlag[uint32]))
	puzzle.StringSliceConverter.Register(FlagFrontend, puzzle.ConvertCallbackFactory1(autoFlag[[]string]))
	puzzle.Float32Converter.Register(FlagFrontend, puzzle.ConvertCallbackFactory1(autoFlag[float32]))
}

func autoFlag[T any](entry *puzzle.Entry[T], fs *flag.FlagSet) error {
	if entry.Metadata.FlagName != "" {
		fs.Var(entry, entry.Metadata.FlagName, entry.Metadata.Description)
	}
	return nil
}
