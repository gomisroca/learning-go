package main

import (
	"fmt"
	"os"
)

// Immediately exit with a given status
func exit() {
	// Defers will not be run when using os.Exit
	defer fmt.Println("!")

	os.Exit(3) // Exit with status 3
}