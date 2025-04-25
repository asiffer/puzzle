package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/asiffer/puzzle"
	"github.com/asiffer/puzzle/flagset"
)

var config = puzzle.NewConfig()

var counter uint64 = 1
var adminUser string = "me"
var secret string = "p4$$w0rD"

func init() {
	puzzle.DefineVar(config, "admin-user", &adminUser)                         // default flag is set to -admin-user
	puzzle.DefineVar(config, "count", &counter, puzzle.WithFlagName("number")) // we redefine it to -number
	puzzle.DefineVar(config, "secret", &secret, puzzle.WithoutFlagName())      // we disable flag for this entry
}

func main() {
	fs, err := flagset.Build(config, "myApp", flag.ContinueOnError)
	if err != nil {
		panic(err)
	}

	// all the config is updated when args are parsed
	if err := fs.Parse(os.Args[1:]); err != nil {
		panic(err)
	}

	fmt.Println(counter, adminUser, secret)
}
