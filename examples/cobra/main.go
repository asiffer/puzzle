package main

import (
	"fmt"

	"github.com/asiffer/puzzle"
	"github.com/asiffer/puzzle/pflagset"
	"github.com/spf13/cobra"
)

var config = puzzle.NewConfig()

var counter uint64 = 1
var adminUser string = "me"
var secret string = "p4$$w0rD"
var verbose = false

var rootCmd = &cobra.Command{
	Use:   "puzzle-cobra",
	Short: "A example of binding puzzle and spf13/cobra",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println(counter, adminUser, secret, verbose)
	},
}

func init() {
	puzzle.DefineVar(config, "admin-user", &adminUser)                           // default flag is set to -admin-user
	puzzle.DefineVar(config, "count", &counter, puzzle.WithFlagName("number"))   // we redefine it to -number
	puzzle.DefineVar(config, "secret", &secret, puzzle.WithoutFlagName())        // we disable flag for this entry
	puzzle.DefineVar(config, "verbose", &verbose, puzzle.WithShortFlagName("v")) // you can use -v
}

func main() {
	pflagset.Populate(config, rootCmd.Flags()) // here is the magic
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
