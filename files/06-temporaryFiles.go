package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// We often want to create temporary files that are deleted when we're done
func temporaryFiles() {
	// CreateTemp creates and opens it
	// First arg is the directory to create the file in, second is the prefix
	f, err := os.CreateTemp("", "sample") // "" is the default dir of the OS
	check(err)

	fmt.Println("Temp file name:", f.Name()) // Name is prefix + random string

	defer os.Remove(f.Name()) // Clean up after ourselves

	// Write some data to the file
	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	// Can also create a temp directory
	dname, err := os.MkdirTemp("", "sampledir")
	check(err)
	fmt.Println("Temp dir name:", dname)

	defer os.RemoveAll(dname) // Clean up after ourselves

	// Now we can create tmp files in our tmp dir with Join
	fname := filepath.Join(dname, "file1")
	err = os.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
}