package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// Go has built-in fns for working with directories in the fs
func directories() {
	err := os.Mkdir("subdir", 0755) // Create a new subdirectory
	check(err)

	defer os.RemoveAll("subdir") // Clean up after ourselves

	// Helper fn to create a new empty file
	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}

	createEmptyFile("subdir/file1")

	// We can create a hierarchy of directories w/ MkdirAll
	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyFile("subdir/parent/file2")
    createEmptyFile("subdir/parent/file3")
    createEmptyFile("subdir/parent/child/file4")

	// ReadDir lists the contents of a directory
	c, err := os.ReadDir("subdir/parent")
	check(err)

	fmt.Println("Listing subdir/parent")
    for _, entry := range c {
        fmt.Println(" ", entry.Name(), entry.IsDir())
    }

	// Chdir changes the current working directory, similar to cd
	err = os.Chdir("subdir/parent/child")
	check(err)
	// Now if we ReadDir, we'll see the contents of subdir/parent/child
	c, err = os.ReadDir(".")
	check(err)

	fmt.Println("Listing subdir/parent/child")
    for _, entry := range c {
        fmt.Println(" ", entry.Name(), entry.IsDir())
    }

	// Now we cd back up to the parent directory
	err = os.Chdir("../../..")
    check(err)

	// We can also visit a dir recursively
	// WalkDir accepts a callback function that is called for each entry
	fmt.Println("Visiting subdir")
    err = filepath.WalkDir("subdir", visit)
}

// visit will be called for each file or directory found during filepath.WalkDir
func visit(path string, d fs.DirEntry, err error) error {
    if err != nil {
        return err
    }
    fmt.Println(" ", path, d.IsDir())
    return nil
}