package main

import (
	"embed"
	"fmt"
)

// the comment go:embed is a compiler directive
// it allows us to embed files into the Go binary at build time

//go:embed folder/single_file.txt
var fileString string

//go:embed folder/single_file.txt
var fileBytes []byte

//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

func embedDirective() {
	// print out the contents of the file
	fmt.Println(fileString)
	fmt.Println(string(fileBytes))

	// Retrieve some files from the embedded folder
	content1, _ := folder.ReadFile("folder/file1.hash")
	fmt.Println(string(content1))

	content2, _ := folder.ReadFile("folder/file2.hash")
	fmt.Println(string(content2))
}