package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/asiffer/puzzle"
	"github.com/asiffer/puzzle/urfave3"
	"github.com/urfave/cli/v3"
)

var config = puzzle.NewConfig()

var counter uint64 = 0
var adminUser string = "me"
var secret string = "p4$$w0rD"

func init() {
	puzzle.DefineVar(config, "admin-user", &adminUser, puzzle.WithShortFlagName("a"))
	puzzle.DefineVar(config, "count", &counter)
	puzzle.DefineVar(config, "secret", &secret)
}

func main() {
	flags0, err := urfave3.Build(config)
	if err != nil {
		panic(err)
	}

	cmd := &cli.Command{
		Name:  "puzzle-urfave3",
		Usage: "A example of binding puzzle and urfave/cli (v3)",
		Flags: flags0,
		Action: func(ctx context.Context, cmd *cli.Command) error {
			fmt.Println(counter, adminUser, secret)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
