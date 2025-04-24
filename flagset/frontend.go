package flagset

import (
	"flag"

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
