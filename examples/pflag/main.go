package main

import (
	"fmt"
	"os"

	"github.com/asiffer/puzzle"
	"github.com/asiffer/puzzle/pflagset"
	"github.com/spf13/pflag"
)

var config = puzzle.NewConfig()

var counter uint64 = 1
var adminUser string = "me"
var secret string = "p4$$w0rD"
var verbose = false

func init() {
	puzzle.DefineVar(config, "admin-user", &adminUser)                           // default flag is set to -admin-user
	puzzle.DefineVar(config, "count", &counter, puzzle.WithFlagName("number"))   // we redefine it to -number
	puzzle.DefineVar(config, "secret", &secret, puzzle.WithoutFlagName())        // we disable flag for this entry
	puzzle.DefineVar(config, "verbose", &verbose, puzzle.WithShortFlagName("v")) // you can use -v
}

func main() {
	fs, err := pflagset.Build(config, "myApp", pflag.ContinueOnError)
	if err != nil {
		panic(err)
	}

	// all the config is updated when args are parsed
	if err := fs.Parse(os.Args[1:]); err != nil {
		panic(err)
	}

	fmt.Println(counter, adminUser, secret, verbose)
}
