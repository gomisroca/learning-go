package main

import (
	"bufio"
	"fmt"
	"os"
)

func writingFiles() {
	// Dump a string (or bytes) into a file
	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	// Again, we probably want more granular control of what we write
	// To do this, we start by opening a file for writing
	f, err := os.Create("/tmp/dat2")
	check(err)

	defer f.Close() // Close the file when we're done

	// We can write byte slices
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// Or strings
	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync() // Flush the file to disk

	// We can also have buffered writers 
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush() // Ensure all buffered data is has been written
}