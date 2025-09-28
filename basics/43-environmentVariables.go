package main

import (
	"fmt"
	"os"
	"strings"
)

func environmentVariables() {
	// Set env
	os.Setenv("FOO", "1")
	// Get env
	fmt.Println("FOO:", os.Getenv("FOO"))
	// If empty, returns the default value
    fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println()
	// List all key/value pairs in the env
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}