package pflagset

import (
	"fmt"

	"github.com/asiffer/puzzle"
	"github.com/spf13/pflag"
)

const PFlagFrontend puzzle.Frontend = "pflagset"

func Build(c *puzzle.Config, name string, h pflag.ErrorHandling) (*pflag.FlagSet, error) {
	flagset := pflag.NewFlagSet(name, h)
	return flagset, Populate(c, flagset)
}

func Populate(c *puzzle.Config, flagset *pflag.FlagSet) error {
	for entry := range c.Entries() {
		if err := entry.Convert(PFlagFrontend, flagset); err != nil {
			return err
		}
	}
	return nil
}

func ToFlags(c *puzzle.Config, useShort bool) []string {
	out := make([]string, 0)
	for e := range c.Entries() {
		fn := e.GetMetadata().FlagName
		sfn := e.GetMetadata().ShortFlagName
		if fn == "" {
			continue
		}

		if useShort && sfn != "" {
			switch b := e.GetValue().(type) {
			case bool:
				if b {
					out = append(out, "-"+sfn)
				} else {
					// falbback to long flag
					out = append(out, fmt.Sprintf("--%s=%s", fn, e.String()))
				}
			default:
				out = append(out, "-"+sfn, e.String())
			}
		} else {
			out = append(out, fmt.Sprintf("--%s=%s", fn, e.String()))
		}
	}
	return out
}
