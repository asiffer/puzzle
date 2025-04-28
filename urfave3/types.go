package urfave3

import (
	"net"
	"time"

	"github.com/asiffer/puzzle"
	"github.com/urfave/cli/v3"
)

func init() {
	// default
	puzzle.BoolConverter.Register(Urfave3Frontend, puzzle.ConvertCallbackFactory2(boolFlag))
	puzzle.DurationConverter.Register(Urfave3Frontend, puzzle.ConvertCallbackFactory2(autoFlag[time.Duration]))
	puzzle.Float64Converter.Register(Urfave3Frontend, puzzle.ConvertCallbackFactory2(autoFlag[float64]))
	puzzle.IntConverter.Register(Urfave3Frontend, puzzle.ConvertCallbackFactory2(autoFlag[int]))
	puzzle.Int64Converter.Register(Urfave3Frontend, puzzle.ConvertCallbackFactory2(autoFlag[int64]))
	puzzle.StringConverter.Register(Urfave3Frontend, puzzle.ConvertCallbackFactory2(autoFlag[string]))
	puzzle.UintConverter.Register(Urfave3Frontend, puzzle.ConvertCallbackFactory2(autoFlag[uint]))
	puzzle.Uint64Converter.Register(Urfave3Frontend, puzzle.ConvertCallbackFactory2(autoFlag[uint64]))
	puzzle.BytesConverter.Register(Urfave3Frontend, puzzle.ConvertCallbackFactory2(autoFlag[[]byte]))
	puzzle.IPConverter.Register(Urfave3Frontend, puzzle.ConvertCallbackFactory2(autoFlag[net.IP]))
	puzzle.Int8Converter.Register(Urfave3Frontend, puzzle.ConvertCallbackFactory2(autoFlag[int8]))
	puzzle.Int16Converter.Register(Urfave3Frontend, puzzle.ConvertCallbackFactory2(autoFlag[int16]))
	puzzle.Int32Converter.Register(Urfave3Frontend, puzzle.ConvertCallbackFactory2(autoFlag[int32]))
	puzzle.Uint8Converter.Register(Urfave3Frontend, puzzle.ConvertCallbackFactory2(autoFlag[uint8]))
	puzzle.Uint16Converter.Register(Urfave3Frontend, puzzle.ConvertCallbackFactory2(autoFlag[uint16]))
	puzzle.Uint32Converter.Register(Urfave3Frontend, puzzle.ConvertCallbackFactory2(autoFlag[uint32]))
	puzzle.StringSliceConverter.Register(Urfave3Frontend, puzzle.ConvertCallbackFactory2(autoFlag[[]string]))
	puzzle.Float32Converter.Register(Urfave3Frontend, puzzle.ConvertCallbackFactory2(autoFlag[float32]))
}

func boolFlag(entry *puzzle.Entry[bool], flags *[]cli.Flag, options []FlagBaseSubsetOption) error {
	f := cli.BoolFlag{
		Name:             entry.Metadata.FlagName,
		Destination:      entry.ValueP,
		Value:            entry.Value,
		Usage:            entry.Metadata.Description,
		Hidden:           false,                       // display flag by default
		Required:         false,                       // flag is not required by default
		TakesFile:        entry.Metadata.IsConfigFile, // flag does not take a file by default
		OnlyOnce:         true,                        // flag can be set only once by default
		ValidateDefaults: true,                        // validate defaults by default
		Local:            false,
		Aliases:          make([]string, 0),
	}
	if entry.Metadata.ShortFlagName != "" {
		f.Aliases = append(f.Aliases, entry.Metadata.ShortFlagName)
	}
	*flags = append(*flags, &f)
	return nil
}

func autoFlag[T any](entry *puzzle.Entry[T], flags *[]cli.Flag, options []FlagBaseSubsetOption) error {
	if entry.Metadata.FlagName == "" {
		return nil
	}
	fb := defaultFlagBase(entry)
	// mofify some internal urfave3 stuff
	fbs := exposeFlagBaseSubset(fb)
	for _, opt := range options {
		opt(fbs)
	}
	*flags = append(*flags, fb)
	return nil
}
