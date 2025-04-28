package pflagset

import (
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
