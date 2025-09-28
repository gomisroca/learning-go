package main

import (
	"flag"
	"fmt"
	"os"
)

// Many command-line tools have subcommands, each with their own flags
// f.e: go get, go build, go run, go test, go fmt, etc.
// With flag we can define subcommands with their own flags
func commandLineSubcommands() {
	// Declare a subcommand
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	// Declare flags for the subcommand
	fooEnable := fooCmd.Bool("enable", false, "enable")
    fooName := fooCmd.String("name", "", "name")

	// Another subcommand with its own flags
	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
    barLevel := barCmd.Int("level", 0, "level")

	// Expect the subcommand as first arg of the program
	if len(os.Args) < 2 {
        fmt.Println("expected 'foo' or 'bar' subcommands")
        os.Exit(1)
    }

	// Check which subcommand is being used
	// Parse the subcommand flags
	switch os.Args[1] {
	case "foo":
        fooCmd.Parse(os.Args[2:])
        fmt.Println("subcommand 'foo'")
        fmt.Println("  enable:", *fooEnable)
        fmt.Println("  name:", *fooName)
        fmt.Println("  tail:", fooCmd.Args())
    case "bar":
        barCmd.Parse(os.Args[2:])
        fmt.Println("subcommand 'bar'")
        fmt.Println("  level:", *barLevel)
        fmt.Println("  tail:", barCmd.Args())
    default:
        fmt.Println("expected 'foo' or 'bar' subcommands")
        os.Exit(1)
    }
}