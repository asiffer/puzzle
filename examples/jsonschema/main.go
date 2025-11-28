package main

import (
	"encoding/json"
	"fmt"

	"github.com/asiffer/puzzle"
	"github.com/asiffer/puzzle/jsonschema"
)

var config = puzzle.NewConfig()

var counter uint64 = 0
var adminUser string = "me"
var secret string = "p4$$w0rD"

func init() {
	puzzle.DefineConfigFile(config, "config", []string{"config.json"})
	puzzle.DefineVar(config, "admin-user", &adminUser) // we directly use the key to read the json
	puzzle.DefineVar(config, "count", &counter)
	puzzle.DefineVar(config, "secret", &secret)
}

func main() {
	// generate the schema from your config
	schema, err := jsonschema.Generate(config)
	if err != nil {
		// /!\ if it fails during parsing the json file the config can be corrupted
		// (only some values are updated)
		panic(err)
	}
	// export the schema
	bytes, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}
