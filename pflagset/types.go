package pflagset

import (
	"net"
	"time"

	"github.com/asiffer/puzzle"
	"github.com/spf13/pflag"
)

func init() {
	// default
	puzzle.BoolConverter.Register(PFlagFrontend, puzzle.ConvertCallbackFactory1(ignoreFlag(boolFlag)))
	puzzle.DurationConverter.Register(PFlagFrontend, puzzle.ConvertCallbackFactory1(ignoreFlag(durationFlag)))
	puzzle.Float64Converter.Register(PFlagFrontend, puzzle.ConvertCallbackFactory1(ignoreFlag(float64Flag)))
	puzzle.IntConverter.Register(PFlagFrontend, puzzle.ConvertCallbackFactory1(ignoreFlag(intFlag)))
	puzzle.Int64Converter.Register(PFlagFrontend, puzzle.ConvertCallbackFactory1(ignoreFlag(int64Flag)))
	puzzle.StringConverter.Register(PFlagFrontend, puzzle.ConvertCallbackFactory1(ignoreFlag(stringFlag)))
	puzzle.UintConverter.Register(PFlagFrontend, puzzle.ConvertCallbackFactory1(ignoreFlag(uintFlag)))
	puzzle.Uint64Converter.Register(PFlagFrontend, puzzle.ConvertCallbackFactory1(ignoreFlag(uint64Flag)))
	puzzle.BytesConverter.Register(PFlagFrontend, puzzle.ConvertCallbackFactory1(ignoreFlag(bytesFlag)))
	puzzle.IPConverter.Register(PFlagFrontend, puzzle.ConvertCallbackFactory1(ignoreFlag(ipFlag)))
	puzzle.Int8Converter.Register(PFlagFrontend, puzzle.ConvertCallbackFactory1(ignoreFlag(int8Flag)))
	puzzle.Int16Converter.Register(PFlagFrontend, puzzle.ConvertCallbackFactory1(ignoreFlag(int16Flag)))
	puzzle.Int32Converter.Register(PFlagFrontend, puzzle.ConvertCallbackFactory1(ignoreFlag(int32Flag)))
	puzzle.Uint8Converter.Register(PFlagFrontend, puzzle.ConvertCallbackFactory1(ignoreFlag(uint8Flag)))
	puzzle.Uint16Converter.Register(PFlagFrontend, puzzle.ConvertCallbackFactory1(ignoreFlag(uint16Flag)))
	puzzle.Uint32Converter.Register(PFlagFrontend, puzzle.ConvertCallbackFactory1(ignoreFlag(uint32Flag)))
	puzzle.StringSliceConverter.Register(PFlagFrontend, puzzle.ConvertCallbackFactory1(ignoreFlag(stringSliceFlag)))
	puzzle.Float32Converter.Register(PFlagFrontend, puzzle.ConvertCallbackFactory1(ignoreFlag(float32Flag)))
}

// supported types

type flagFun[T any] func(entry *puzzle.Entry[T], fs *pflag.FlagSet) error

func ignoreFlag[T any](f flagFun[T]) flagFun[T] {
	return func(entry *puzzle.Entry[T], fs *pflag.FlagSet) error {
		if entry.Metadata.FlagName == "" {
			return nil
		}
		return f(entry, fs)
	}
}

func boolFlag(entry *puzzle.Entry[bool], fs *pflag.FlagSet) error {
	if entry.Metadata.ShortFlagName == "" {
		fs.BoolVar(entry.ValueP, entry.Metadata.FlagName, entry.Value, entry.Metadata.Description)
	} else {
		fs.BoolVarP(entry.ValueP, entry.Metadata.FlagName, entry.Metadata.ShortFlagName, entry.Value, entry.Metadata.Description)
	}
	return nil
}

func durationFlag(entry *puzzle.Entry[time.Duration], fs *pflag.FlagSet) error {
	if entry.Metadata.ShortFlagName == "" {
		fs.DurationVar(entry.ValueP, entry.Metadata.FlagName, entry.Value, entry.Metadata.Description)
	} else {
		fs.DurationVarP(entry.ValueP, entry.Metadata.FlagName, entry.Metadata.ShortFlagName, entry.Value, entry.Metadata.Description)
	}
	return nil
}

func float32Flag(entry *puzzle.Entry[float32], fs *pflag.FlagSet) error {
	if entry.Metadata.ShortFlagName == "" {
		fs.Float32Var(entry.ValueP, entry.Metadata.FlagName, entry.Value, entry.Metadata.Description)
	} else {
		fs.Float32VarP(entry.ValueP, entry.Metadata.FlagName, entry.Metadata.ShortFlagName, entry.Value, entry.Metadata.Description)
	}
	return nil
}

func float64Flag(entry *puzzle.Entry[float64], fs *pflag.FlagSet) error {
	if entry.Metadata.ShortFlagName == "" {
		fs.Float64Var(entry.ValueP, entry.Metadata.FlagName, entry.Value, entry.Metadata.Description)
	} else {
		fs.Float64VarP(entry.ValueP, entry.Metadata.FlagName, entry.Metadata.ShortFlagName, entry.Value, entry.Metadata.Description)
	}
	return nil
}

func intFlag(entry *puzzle.Entry[int], fs *pflag.FlagSet) error {
	if entry.Metadata.ShortFlagName == "" {
		fs.IntVar(entry.ValueP, entry.Metadata.FlagName, entry.Value, entry.Metadata.Description)
	} else {
		fs.IntVarP(entry.ValueP, entry.Metadata.FlagName, entry.Metadata.ShortFlagName, entry.Value, entry.Metadata.Description)
	}
	return nil
}

func int8Flag(entry *puzzle.Entry[int8], fs *pflag.FlagSet) error {
	if entry.Metadata.ShortFlagName == "" {
		fs.Int8Var(entry.ValueP, entry.Metadata.FlagName, entry.Value, entry.Metadata.Description)
	} else {
		fs.Int8VarP(entry.ValueP, entry.Metadata.FlagName, entry.Metadata.ShortFlagName, entry.Value, entry.Metadata.Description)
	}
	return nil
}

func int16Flag(entry *puzzle.Entry[int16], fs *pflag.FlagSet) error {
	if entry.Metadata.ShortFlagName == "" {
		fs.Int16Var(entry.ValueP, entry.Metadata.FlagName, entry.Value, entry.Metadata.Description)
	} else {
		fs.Int16VarP(entry.ValueP, entry.Metadata.FlagName, entry.Metadata.ShortFlagName, entry.Value, entry.Metadata.Description)
	}
	return nil
}

func int32Flag(entry *puzzle.Entry[int32], fs *pflag.FlagSet) error {
	if entry.Metadata.ShortFlagName == "" {
		fs.Int32Var(entry.ValueP, entry.Metadata.FlagName, entry.Value, entry.Metadata.Description)
	} else {
		fs.Int32VarP(entry.ValueP, entry.Metadata.FlagName, entry.Metadata.ShortFlagName, entry.Value, entry.Metadata.Description)
	}
	return nil
}

func int64Flag(entry *puzzle.Entry[int64], fs *pflag.FlagSet) error {
	if entry.Metadata.ShortFlagName == "" {
		fs.Int64Var(entry.ValueP, entry.Metadata.FlagName, entry.Value, entry.Metadata.Description)
	} else {
		fs.Int64VarP(entry.ValueP, entry.Metadata.FlagName, entry.Metadata.ShortFlagName, entry.Value, entry.Metadata.Description)
	}
	return nil
}

func uintFlag(entry *puzzle.Entry[uint], fs *pflag.FlagSet) error {
	if entry.Metadata.ShortFlagName == "" {
		fs.UintVar(entry.ValueP, entry.Metadata.FlagName, entry.Value, entry.Metadata.Description)
	} else {
		fs.UintVarP(entry.ValueP, entry.Metadata.FlagName, entry.Metadata.ShortFlagName, entry.Value, entry.Metadata.Description)
	}
	return nil
}

func uint8Flag(entry *puzzle.Entry[uint8], fs *pflag.FlagSet) error {
	if entry.Metadata.ShortFlagName == "" {
		fs.Uint8Var(entry.ValueP, entry.Metadata.FlagName, entry.Value, entry.Metadata.Description)
	} else {
		fs.Uint8VarP(entry.ValueP, entry.Metadata.FlagName, entry.Metadata.ShortFlagName, entry.Value, entry.Metadata.Description)
	}
	return nil
}

func uint16Flag(entry *puzzle.Entry[uint16], fs *pflag.FlagSet) error {
	if entry.Metadata.ShortFlagName == "" {
		fs.Uint16Var(entry.ValueP, entry.Metadata.FlagName, entry.Value, entry.Metadata.Description)
	} else {
		fs.Uint16VarP(entry.ValueP, entry.Metadata.FlagName, entry.Metadata.ShortFlagName, entry.Value, entry.Metadata.Description)
	}
	return nil
}

func uint32Flag(entry *puzzle.Entry[uint32], fs *pflag.FlagSet) error {
	if entry.Metadata.ShortFlagName == "" {
		fs.Uint32Var(entry.ValueP, entry.Metadata.FlagName, entry.Value, entry.Metadata.Description)
	} else {
		fs.Uint32VarP(entry.ValueP, entry.Metadata.FlagName, entry.Metadata.ShortFlagName, entry.Value, entry.Metadata.Description)
	}
	return nil
}

func uint64Flag(entry *puzzle.Entry[uint64], fs *pflag.FlagSet) error {
	if entry.Metadata.ShortFlagName == "" {
		fs.Uint64Var(entry.ValueP, entry.Metadata.FlagName, entry.Value, entry.Metadata.Description)
	} else {
		fs.Uint64VarP(entry.ValueP, entry.Metadata.FlagName, entry.Metadata.ShortFlagName, entry.Value, entry.Metadata.Description)
	}
	return nil
}

func stringFlag(entry *puzzle.Entry[string], fs *pflag.FlagSet) error {
	if entry.Metadata.ShortFlagName == "" {
		fs.StringVar(entry.ValueP, entry.Metadata.FlagName, entry.Value, entry.Metadata.Description)
	} else {
		fs.StringVarP(entry.ValueP, entry.Metadata.FlagName, entry.Metadata.ShortFlagName, entry.Value, entry.Metadata.Description)
	}
	return nil
}

func bytesFlag(entry *puzzle.Entry[[]byte], fs *pflag.FlagSet) error {
	if entry.Metadata.ShortFlagName == "" {
		switch entry.Metadata.Format {
		case "base64":
			fs.BytesBase64Var(entry.ValueP, entry.Metadata.FlagName, entry.Value, entry.Metadata.Description)
		case "base32":
			fs.Var(entry, entry.Metadata.FlagName, entry.Metadata.Description)
		default:
			fs.BytesHexVar(entry.ValueP, entry.Metadata.FlagName, entry.Value, entry.Metadata.Description)
		}
	} else {
		switch entry.Metadata.Format {
		case "base64":
			fs.BytesBase64VarP(entry.ValueP, entry.Metadata.FlagName, entry.Metadata.ShortFlagName, entry.Value, entry.Metadata.Description)
		case "base32":
			fs.VarP(entry, entry.Metadata.FlagName, entry.Metadata.ShortFlagName, entry.Metadata.Description)
		default:
			fs.BytesHexVarP(entry.ValueP, entry.Metadata.FlagName, entry.Metadata.ShortFlagName, entry.Value, entry.Metadata.Description)
		}
	}
	return nil
}

func ipFlag(entry *puzzle.Entry[net.IP], fs *pflag.FlagSet) error {
	if entry.Metadata.ShortFlagName == "" {
		fs.IPVar(entry.ValueP, entry.Metadata.FlagName, entry.Value, entry.Metadata.Description)
	} else {
		fs.IPVarP(entry.ValueP, entry.Metadata.FlagName, entry.Metadata.ShortFlagName, entry.Value, entry.Metadata.Description)
	}
	return nil
}

func stringSliceFlag(entry *puzzle.Entry[[]string], fs *pflag.FlagSet) error {
	if entry.Metadata.ShortFlagName == "" {
		fs.StringSliceVar(entry.ValueP, entry.Metadata.FlagName, entry.Value, entry.Metadata.Description)
	} else {
		fs.StringSliceVarP(entry.ValueP, entry.Metadata.FlagName, entry.Metadata.ShortFlagName, entry.Value, entry.Metadata.Description)
	}
	return nil
}
