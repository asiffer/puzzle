package main

import (
	"fmt"

	"github.com/asiffer/puzzle"
)

var config = puzzle.NewConfig()

var counter uint64 = 0
var adminUser string = "me"
var secret string = "p4$$w0rD"

func init() {
	puzzle.DefineVar(config, "admin-user", &adminUser)                   // default env name is set to ADMIN_USER
	puzzle.DefineVar(config, "count", &counter, puzzle.WithEnvName("N")) // we redefine it to N
	puzzle.DefineVar(config, "secret", &secret, puzzle.WithoutEnv())     // we disable env for this entry
}

func main() {
	// update the config from env
	if err := puzzle.ReadEnv(config); err != nil {
		panic(err)
	}
	fmt.Println(counter, adminUser, secret)
}
