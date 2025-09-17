package main

import (
	"fmt"
	"os"
)

// Defer is used to ensure a fn call is performed later, similar to finally in js
func advancedDefers() {
	f := createFile("/tmp/defer.txt")
	defer closeFile(f) // This will be executed when the function returns
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	// Important to check for errors in a closing fn, even if deferred
	if err != nil {
		panic(err)
	}
}