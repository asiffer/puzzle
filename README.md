<p align="center">
    <img src="https://raw.githubusercontent.com/asiffer/puzzle/master/assets/logo.svg" width="320" alt="asiffer/puzzle">
</p>

<p align="center">
    <b>puzzle - Bottom-up golang configuration library</b>
</p>


## Introduction

**All we need is configuration**. Yes we also need commands, flags, env variables, config files... but they are just *frontends*. Why should we create configuration from command flags? And not the opposite?

`puzzle` aims to centralize the configuration management and to **automatically create the bindings** you need from other sources at runtime (like environment, json files, [`flag`](https://pkg.go.dev/flag), [`spf13/cobra`](https://github.com/spf13/cobra), [`urfave/cli`](https://github.com/urfave/cli)...). No annotations, just generics.

[![Go Report Card](https://goreportcard.com/badge/github.com/asiffer/puzzle)](https://goreportcard.com/report/github.com/asiffer/puzzle) 
[![Test](https://github.com/asiffer/puzzle/actions/workflows/test.yml/badge.svg)](https://github.com/asiffer/puzzle/actions/workflows/test.yml) 
[![Go Reference](https://pkg.go.dev/badge/github.com/asiffer/puzzle.svg)](https://pkg.go.dev/github.com/asiffer/puzzle)
![GitHub License](https://img.shields.io/github/license/asiffer/puzzle)
![Library base size](https://img.shields.io/badge/size-435_KB-green)
[![codecov](https://codecov.io/github/asiffer/puzzle/graph/badge.svg?token=7R7S74GVRF)](https://codecov.io/github/asiffer/puzzle)

## Install

```shell
go get -u github.com/asiffer/puzzle
```

```go
import "github.com/asiffer/puzzle"
```

## Get started

First define a `Config` object.
```go
// config.go

var config = puzzle.NewConfig()
```

Then, anywhere in your code, define configuration variables.

```go
func init() {
    puzzle.Define[string](config, "question", "The Ultimate Question of Life, the Universe and Everything")
    puzzle.Define[int](config, "answer", 42)
}
```

You can also be responsible of variable storage.

```go
import "github.com/asiffer/puzzle"

var question string = "The Ultimate Question of Life, the Universe and Everything"
var anwser int = 42

func init() {
    puzzle.DefineVar[string](config, "question", &question)
    puzzle.DefineVar[int](config, "answer", &anwser)
}
```

You can access it the `Get` method.

```go
func main() {
    q, err := puzzle.Get[string](config, "question")
    if err != nil {
        panic(err)
    }
    fmt.Println(q == question)
}
```

But, the most interesting thing is that you can directly create the related flags.

```go
import "github.com/asiffer/puzzle/flagset"

func main() {
    // from your config, generate the flagset
    fs, err := flagset.Build(config, "myApp", flag.PanicOnError)
    if err != nil {
        panic(err)
    }
    // fill the config while parsing flags
    if err := fs.Parse(os.Args); err != nil {
        panic(err)
    }
    
    // access via puzzle
    q, err := puzzle.Get[string](config, "question")
    if err != nil {
        panic(err)
    } else {
        fmt.Println(q)
    }
    // or directly (if defined with puzzle.DefineVar[string](config, "question", &question))
    fmt.Println(question)
}
```


## Frontends

Once a config is defined. The goal of **puzzle** is to be able to automatically binds to incoming source (a.k.a. *frontends*). 

### Environment

```go
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

```

### CLI flags

There are several options to parse input flags. 
The **puzzle** library aims to target the most popular.

#### `flag`

```go
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

```



#### `spf13/pflag`

```go
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

```

#### `spf13/cobra` 

```go
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

```


#### `urfave/cli/v3`

```go
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

```


### JSON file

```go
package main

import (
	"fmt"

	"github.com/asiffer/puzzle"
	"github.com/asiffer/puzzle/jsonfile"
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
	if err := jsonfile.ReadJSON(config); err != nil {
		// /!\ if it fails during parsing the json file the config can be corrupted
		// (only some values are updated)
		panic(err)
	}
	// all the config is updated
	fmt.Println(counter, adminUser, secret)
}

```


## Customizations

While defining a config variable you can customize some of its properties, useful for subsequent flag generation or env parsing tasks.

```go
var exampleVar time.Duration = 5 * time.Minute

puzzle.DefineVar[time.Duration](
    config, // your config
    "example_var", // the key to access it
    &exampleVar, // the storage location
    puzzle.WithDescription("my example variable"), // for flag usage notably
    puzzle.WithEnvName("EXAMPLE"), // instead of EXAMPLE_VAR
    puzzle.WithFlagName("example"), // instead of example-var
    puzzle.WithShortFlagName("e"), // no short flag by default (used for pflag)
)
```

## Patterns

### Helpers (DX)

In the case where we have a single config, we can create wrappers to ease config definition. In the example below, we hide `konf` in a dedicated `config` package, exposing only (simpler) i/o functions.

```go
// config/config.go
package config

var konf = puzzle.NewConfig()

func Define[T any](key string, defaultValue T, options ...puzzle.MetadataOption) error {
    return puzzle.Define[T](konf, key, defaultValue, &question, options...)
}

func DefineVar[T any](key string, boundVariable *T, options ...puzzle.MetadataOption) error {
    return puzzle.DefineVar[T](konf, key, boundVariable, &question, options...)
}

func Get[T](key string) (T, error) {
    return puzzle.Get[T](konf, key)
}
```

### Config struct

If your whole config is stored in a struct, you probably need another library to manage it. At small scale, you can use `puzzle` on every attribute.

```go
var config = puzzle.NewConfig()

type ConfigurationType struct {
    Level int 
    Verbose bool 
    Name string
    Modules []string
}

var Configuration = Configuration{
    Level: 1,
    Verbose: false,
    Name: "remote",
    Modules: []string{"user", "auth"}
}

func init() {
    puzzle.DefineVar[int](config, "level", &Configuration.Level)
    puzzle.DefineVar[bool](config, "verbose", &Configuration.Verbose)
    puzzle.DefineVar[string](config, "remote", &Configuration.Name)
    puzzle.DefineVar[[]string](config, "modules", &Configuration.Modules)
}
```

### Config file 

In many cases, you may need to read the config from both the command line and a config file (also provided by the command line). To handle this case, you should split the process (ignoring or considering only the config key).

#### Configurable config file

In this case, the config file can be configured by the end user. The following example gives an example where we take its value from cli flags (but it could be read from env or any other supported source).

```go
var level int = 3

var config = puzzle.NewConfig()

func init() {
    puzzle.DefineVar[int](config, "level", &level)
    puzzle.DefineConfigFile(config, "config", []string{"conf.json", "/etc/app/conf.json"})
}

func main() {
    // if you need to read the config from the command line
    fs, err := flagset.Build(config.Only("config"), "myApp", flag.PanicOnError)
    if err != nil {
        panic(err)
    }

    // set the value of the config file
    if err := fs.Parse(os.Args); err != nil {
        panic(err)
    }
    // here the value of the config file is populated
    // we just have to read it with ReadJSON()
    // (puzzle looks for the value, opens the file and reads it)
    if err := config.ReadJSON(); err != nil {
        panic(err)
    }
    // then we can read other flags
    fs, err = flagset.Build(config.Ignoring("config"), "myApp", flag.PanicOnError)
    if err != nil {
        panic(err)
    }
    // here all the config is then populated
}
```

#### Hardcoded config file

If you don't need to set the config file from command line (only using your default values), it is a bit simpler.


```go
var level int = 3

var config = puzzle.NewConfig()

func init() {
    puzzle.DefineVar[int](config, "level", &level)
    puzzle.DefineConfigFile(config, "config", []string{"conf.json"})
}

func main() {
    // generally we first read the config file
    if err := config.ReadJSON(); err != nil {
        panic(err)
    }
    // and then we override the values with the defined flags
    fs, err := flagset.Build(config.Ignoring("config"), "myApp", flag.PanicOnError)
    if err != nil {
        panic(err)
    }

    if err := fs.Parse(os.Args); err != nil {
        panic(err)
    }
    // here all the config is then populated
}
```



## Supported types

| Type            | Supported          |
| --------------- | ------------------ |
| `bool`          | :white_check_mark: |
| `time.Duration` | :white_check_mark: |
| `float32`       | :white_check_mark: |
| `float64`       | :white_check_mark: |
| `int`           | :white_check_mark: |
| `int8`          | :white_check_mark: |
| `int16`         | :white_check_mark: |
| `int32`         | :white_check_mark: |
| `int64`         | :white_check_mark: |
| `string`        | :white_check_mark: |
| `uint`          | :white_check_mark: |
| `uint8`         | :white_check_mark: |
| `uint16`        | :white_check_mark: |
| `uint32`        | :white_check_mark: |
| `uint64`        | :white_check_mark: |
| `[]byte`        | :white_check_mark: |
| `[]string`      | :white_check_mark: |
| `net.IP`        | :white_check_mark: |

## Developer

This library is built around boilerplate code so all the contributions are welcome! 
Naturally, we may need to support either new types or new frontends.

### Supporting a new type

To support a new types, several steps must be performed. 
First a new `xxxx.go` file must be created at the root of the project where `xxxx` is the type. 

This file must define a converter specific to the new type.

```go
// xxxx.go
package puzzle

var XxxxConverter = newConverter(xxxxConverter)

func xxxxConverter(entry *Entry[xxxx], stringValue string) error {
	value, err := xxxxFromString(stringValue) // here is the paramount step where we must be able to parse it from string
	if err != nil {
		return err
	}
	*entry.ValueP = value
	entry.Value = value
	return nil
}
```

Then this converter must be bound to the related entry. 
It is done in the `wire()` method in `entry.go`.

```go
// Wire performs all the plumbing
func (e *Entry[T]) wire() {
	switch z := any(e).(type) {
	case *Entry[bool]:
		z.converter = BoolConverter
	case *Entry[time.Duration]:
		z.converter = DurationConverter
	// ...
    case *Entry[xxxx]: // <- new entry type
        z.converter = XxxxConverter
    // ...
    }
}
```

Finally, this support should be propagated to all the frontends (look at the sub-packages).


### Supporting a new frontend

To create a new frontend to bind the **puzzle** config to, a new sub-package must be created.

```plaintext
├─ README.md
├─ go.mod
├─ go.sum
├─ ...*.go
├─ zzzzzzz/ <- new folder
│  └─ frontend.go
└─ ...
```