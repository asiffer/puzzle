package flagset

import (
	"flag"

	"github.com/asiffer/puzzle"
)

const FlagFrontend puzzle.Frontend = "flagset"

func Build(c *puzzle.Config, name string, h flag.ErrorHandling) (*flag.FlagSet, error) {
	flagset := flag.NewFlagSet(name, h)
	return flagset, Populate(c, flagset)
}

func Populate(c *puzzle.Config, flagset *flag.FlagSet) error {
	for entry := range c.Entries() {
		if err := entry.Convert(FlagFrontend, flagset); err != nil {
			return err
		}
	}
	return nil
}
