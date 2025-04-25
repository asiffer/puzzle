package flagset

import (
	"flag"
	"fmt"

	"github.com/asiffer/puzzle"
)

const FlagFrontend puzzle.Frontend = "flagset"

func Build(c *puzzle.Config, name string, h flag.ErrorHandling) (*flag.FlagSet, error) {
	flagset := flag.NewFlagSet(name, h)
	for entry := range c.Entries() {
		if err := entry.Convert(FlagFrontend, flagset); err != nil {
			return nil, err
		}
	}
	return flagset, nil
}

func ToFlags(c *puzzle.Config) []string {
	out := make([]string, 0)
	for e := range c.Entries() {
		fn := e.GetMetadata().FlagName
		if fn == "" {
			continue
		}
		out = append(out, fmt.Sprintf("-%s=%s", fn, e.String()))
	}
	return out
}
