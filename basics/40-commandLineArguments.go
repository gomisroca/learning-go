package main

import (
	"fmt"
	"os"
)

// Command line args are a common way to give a program execution parameters
// f.e: go run hello.go uses run and hello.go args for the go program
func commandLineArguments() {
	// os.Args is a slice of the command line args
	// First value is the path to the program
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// We can access the args by index
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}