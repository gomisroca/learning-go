package main

import (
	"fmt"
	"os"
)

type point struct {
    x, y int
}

// Go has very powerful string formatting capabilities.
func stringFormatting() {
	p := point{1, 2}
	// has several printing "verbs" that format general Go values
	fmt.Printf("struct1: %v\n", p) // prints struct
	fmt.Printf("struct2: %+v\n", p) // prints struct w/ field names
	fmt.Printf("struct3: %#v\n", p) // prints struct in Go syntax

	fmt.Printf("type: %T\n", p) // prints type of struct
	fmt.Printf("bool: %t\n", true) // formats bool
	fmt.Printf("int: %d\n", 123) // formats int

	fmt.Printf("bin: %b\n", 14) // formats binary
	fmt.Printf("char: %c\n", 33) // prints the character corresponding to the given integer

	fmt.Printf("hex: %x\n", 456) // provides hex encodimg

	// float formattings
	fmt.Printf("float1: %f\n", 78.9) // basic decimal formatting
	fmt.Printf("float2: %e\n", 123400000.0) // scientific notation formatting
    fmt.Printf("float3: %E\n", 123400000.0) // scientific notation formatting

	// string formattings
	fmt.Printf("str1: %s\n", "\"string\"") // basic string printing
	fmt.Printf("str2: %q\n", "\"string\"") // double quotes
	fmt.Printf("str3: %x\n", "hex this") // base-16

	fmt.Printf("pointer: %p\n", &p) // prints the memory address of the variable

	// width and precision
	fmt.Printf("width1: |%6s|%6s|\n", "foo", "b") // right-aligned
	fmt.Printf("width2: |%-6s|%-6s|\n", "foo", "b") // left-aligned

	s := fmt.Sprintf("sprintf: a %s", "string")
    fmt.Println(s)

	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}