package urfave3

import (
	"net"
	"time"

	"github.com/asiffer/puzzle"
	"github.com/urfave/cli/v3"
)

func init() {
	// default
	puzzle.BoolConverter.Register(Urfave3Frontend, puzzle.ConvertCallbackFactory1(autoFlag[bool]))
	puzzle.DurationConverter.Register(Urfave3Frontend, puzzle.ConvertCallbackFactory1(autoFlag[time.Duration]))
	puzzle.Float64Converter.Register(Urfave3Frontend, puzzle.ConvertCallbackFactory1(autoFlag[float64]))
	puzzle.IntConverter.Register(Urfave3Frontend, puzzle.ConvertCallbackFactory1(autoFlag[int]))
	puzzle.Int64Converter.Register(Urfave3Frontend, puzzle.ConvertCallbackFactory1(autoFlag[int64]))
	puzzle.StringConverter.Register(Urfave3Frontend, puzzle.ConvertCallbackFactory1(autoFlag[string]))
	puzzle.UintConverter.Register(Urfave3Frontend, puzzle.ConvertCallbackFactory1(autoFlag[uint]))
	puzzle.Uint64Converter.Register(Urfave3Frontend, puzzle.ConvertCallbackFactory1(autoFlag[uint64]))
	puzzle.BytesConverter.Register(Urfave3Frontend, puzzle.ConvertCallbackFactory1(autoFlag[[]byte]))
	puzzle.IPConverter.Register(Urfave3Frontend, puzzle.ConvertCallbackFactory1(autoFlag[net.IP]))
	puzzle.Int8Converter.Register(Urfave3Frontend, puzzle.ConvertCallbackFactory1(autoFlag[int8]))
	puzzle.Int16Converter.Register(Urfave3Frontend, puzzle.ConvertCallbackFactory1(autoFlag[int16]))
	puzzle.Int32Converter.Register(Urfave3Frontend, puzzle.ConvertCallbackFactory1(autoFlag[int32]))
	puzzle.Uint8Converter.Register(Urfave3Frontend, puzzle.ConvertCallbackFactory1(autoFlag[uint8]))
	puzzle.Uint16Converter.Register(Urfave3Frontend, puzzle.ConvertCallbackFactory1(autoFlag[uint16]))
	puzzle.Uint32Converter.Register(Urfave3Frontend, puzzle.ConvertCallbackFactory1(autoFlag[uint32]))
	puzzle.StringSliceConverter.Register(Urfave3Frontend, puzzle.ConvertCallbackFactory1(autoFlag[[]string]))
	puzzle.Float32Converter.Register(Urfave3Frontend, puzzle.ConvertCallbackFactory1(autoFlag[float32]))
}

func autoFlag[T any](entry *puzzle.Entry[T], flags *[]cli.Flag) error {
	if entry.Metadata.FlagName == "" {
		return nil
	}
	fb := defaultFlagBase(entry)
	*flags = append(*flags, fb)
	return nil
}
