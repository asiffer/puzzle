package urfave3

import (
	"github.com/asiffer/puzzle"

	"github.com/urfave/cli/v3"
)

const Urfave3Frontend puzzle.Frontend = "urfave3"

type FB[T any] cli.FlagBase[T, interface{}, cli.ValueCreator[T, interface{}]]

type EntryCreator[T any] struct{}

func (ec EntryCreator[T]) Create(v T, p *T, c *puzzle.Entry[T]) cli.Value {
	// directly pass the entry
	return c
}

func (ec EntryCreator[T]) ToString(v T) string {
	e := puzzle.NewEntry[T]("")
	e.Value = v
	e.ValueP = &v
	return e.String()
}

// FlagBaseSubset is a subset of the FlagBase struct that contains only the fields
// outside of the generic ones
type FlagBaseSubset struct {
	Name             *string
	Category         *string
	DefaultText      *string
	HideDefault      *bool
	Usage            *string
	Sources          *cli.ValueSourceChain
	Required         *bool
	Hidden           *bool
	Local            *bool
	Aliases          []string
	TakesFile        *bool
	OnlyOnce         *bool
	ValidateDefaults *bool
}

// func exposeFlagBaseSubset[T any](f *cli.FlagBase[T, interface{}, EntryCreator[T]]) *FlagBaseSubset {
// 	return &FlagBaseSubset{
// 		Name:             &f.Name,
// 		Category:         &f.Category,
// 		DefaultText:      &f.DefaultText,
// 		HideDefault:      &f.HideDefault,
// 		Usage:            &f.Usage,
// 		Sources:          &f.Sources,
// 		Required:         &f.Required,
// 		Hidden:           &f.Hidden,
// 		Local:            &f.Local,
// 		Aliases:          f.Aliases,
// 		TakesFile:        &f.TakesFile,
// 		OnlyOnce:         &f.OnlyOnce,
// 		ValidateDefaults: &f.ValidateDefaults,
// 	}
// }

func defaultFlagBase[T any](entry *puzzle.Entry[T]) *cli.FlagBase[T, *puzzle.Entry[T], EntryCreator[T]] {
	f := cli.FlagBase[T, *puzzle.Entry[T], EntryCreator[T]]{
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
		Config:           entry, // inject our entry as config
		Aliases:          make([]string, 0),
	}
	if entry.Metadata.ShortFlagName != "" {
		f.Aliases = append(f.Aliases, entry.Metadata.ShortFlagName)
	}
	return &f
}

func Build(c *puzzle.Config) ([]cli.Flag, error) {
	flags := make([]cli.Flag, 0)
	return flags, Populate(c, &flags)
}

func Populate(c *puzzle.Config, flags *[]cli.Flag) error {
	for entry := range c.Entries() {
		if err := entry.Convert(Urfave3Frontend, flags); err != nil {
			return err
		}
	}
	return nil
}
