package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// A line filter is a type of program that reads input on stdin,
// processes it, and prints a derivated result to stdout.
// gred and sed are two examples of line filters.

func lineFilters() {
	// Wrap the unbuffered os.Stdin w/ a buffered scanner
	// This lets us use the Scan method, which reads a line
	scanner := bufio.NewScanner(os.Stdin)

	// Scan the input line by line
	for scanner.Scan() {
		// Convert the line to uppercase
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}

	// Check for errors during Scan
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}