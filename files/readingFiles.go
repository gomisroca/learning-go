package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Reading files requires checking most calls for errors
// This helper streamlines the process
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func readingFiles() {
	// Commit a file's entire content into memory
	dat, err := os.ReadFile("/tmp/dat")
    check(err)
    fmt.Print(string(dat))

    // But we often want more control over which parts of the file we read
    // To do this, start by opening the file
    f, err := os.Open("/tmp/dat")
    check(err)

    // Read some bytes from the beginning of the file
    // in this case, 5 bytes
    b1 := make([]byte, 5)
    n1, err := f.Read(b1)
    check(err)
    fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

    // We can also Seek to a known position in the file and read from there
    o2, err := f.Seek(6, io.SeekStart)
    check(err)
    b2 := make([]byte, 2)
    n2, err := f.Read(b2)
    check(err)
    fmt.Printf("%d bytes @ %d\n", n2, o2)
    fmt.Printf("%v\n", string(b2[:n2]))
    
    // Seek to current cursor position
    _, err = f.Seek(2, io.SeekCurrent)
    check(err)

    // Seek from end of file
    _, err = f.Seek(-4, io.SeekEnd)
    check(err)

    // Seek from start of file
    o3, err := f.Seek(6, io.SeekStart)
    check(err)
    b3 := make([]byte, 2)
    n3, err := io.ReadAtLeast(f, b3, 2)
    check(err)
    fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

    // "Rewind" to the start of the file
    _, err = f.Seek(0, io.SeekStart)
    check(err)

    // Buffered reader
    r4 := bufio.NewReader(f)
    b4, err := r4.Peek(5)
    check(err)
    fmt.Printf("5 bytes: %s\n", string(b4))

    // Close the file
    f.Close()
}