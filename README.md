Sure! Below is a README for your package.

---

# Arguments

`arguments` is a Go package that simplifies the process of parsing command-line arguments. It supports both short (`-a`) and long (`--argument`) flags, and can automatically generate help messages.

## Installation

To install the package, run:
```sh
go get github.com/coalaura/arguments
```

## Usage

Here's a quick example of how to use the `arguments` package in your Go project.

### Registering Arguments

First, you need to register the arguments you want to parse. You can register both short and long arguments and provide default values for them.

```go
package main

import (
    "github.com/coalaura/arguments"
)

func main() {
    var (
        full    string
        short   string
        long    string
        num     uint32
        boolean bool
    )

    arguments.Register("string", 's', &str).WithHelp("A string argument with both a short and long name")
    arguments.Register("", 'a', &short).WithHelp("A short argument with no long name")
    arguments.Register("long", 0, &long).WithHelp("A long argument with no short name")
    arguments.Register("number", 'n', &num).WithHelp("A number argument")
    arguments.Register("boolean", 'b', &boolean).WithHelp("A boolean argument")

    arguments.MustParse()
}
```

### Parsing Arguments

After registering the arguments, you need to call `MustParse()` to parse the command-line arguments. You can also call `Parse()` if you want to handle errors yourself.

```go
arguments.MustParse()
```

### Help

You can display a help message that lists all the registered arguments and their descriptions by calling `ShowHelp()` or `ShowHelpAndExit()`.

```go
arguments.ShowHelp(true)
```

Alternatively you can call `RegisterHelp()` to register a help argument, which will display the help message and exit the program if set.

```go
arguments.RegisterHelp(true)
```

## Example

Below is a complete example showing how to use the package:

```go
package main

import (
    "github.com/coalaura/arguments"
    "fmt"
)

func main() {
    var (
        input    string
        output   string
        number   uint32
        boolean  bool
    )

    arguments.Register("input", 'i', &input).WithHelp("The input file")
    arguments.Register("output", 0, &output).WithHelp("The output file")
    arguments.Register("", 'n', &number).WithHelp("Number of things")
    arguments.Register("bool", 'b', &boolean).WithHelp("A boolean flag")

    arguments.MustParse()

    fmt.Printf("Input: %s\n", input)
    fmt.Printf("Output: %s\n", output)
    fmt.Printf("Number: %d\n", number)
    fmt.Printf("Boolean: %v\n", boolean)

    arguments.ShowHelp(true)
}
```

## Tests

To run the tests for this package, use the `go test` command:

```sh
go test
```
