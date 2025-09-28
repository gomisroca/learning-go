package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

// filepath provides a set of functions to parse and construct file paths.
func filePaths() {
	// Join is used to construct paths
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p: ", p)

	// Should always use Join instead of adding / or \ manually
	fmt.Println(filepath.Join("dir1//", "filename"))
    fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	// Dir and Base split a path to the dir or file.
	fmt.Println("Dir(p):", filepath.Dir(p))
    fmt.Println("Base(p):", filepath.Base(p))

	// We can check if a pth is absolute
	fmt.Println(filepath.IsAbs("dir/file"))
    fmt.Println(filepath.IsAbs("/dir/file"))

	// We can get the exntension of a file
	filename := "config.json"
	ext := filepath.Ext(filename)
	fmt.Println(ext)

	// We can get the name of a file without the extension
	fmt.Println(strings.TrimSuffix(filename, ext))

	// Rel fins the relative path between a base and a target
	// It will return an error if the target cannot be reached from the base
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}