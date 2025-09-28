package main

import (
	"flag"
	"fmt"
)

// Command line flags are a common way to specify program options
// f.e: wc -l uses -l flag for the wc program

// Go provides a flag pkg for basic flag parsing
func commandLineFlags() {
	// We can declare str, int and bool flags

	// Here we declare a string flag
	// its name is "word", 
	// its default value is "foo", 
	// and its description is "a string"
	wordPtr := flag.String("word", "foo", "a string")

	// The two other types are declared similarly
	numbPtr := flag.Int("numb", 42, "an int")
    forkPtr := flag.Bool("fork", false, "a bool")

	// We can declare an option that uses an existing var
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var") // we pass the pointer to the var

	flag.Parse() // Parse the command line args

	fmt.Println("word:", *wordPtr)
    fmt.Println("numb:", *numbPtr)
    fmt.Println("fork:", *forkPtr)
    fmt.Println("svar:", svar)
    fmt.Println("tail:", flag.Args())
}