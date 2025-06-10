package main

import (
	"fmt"
	"io"
	"strings"
)

type InfiniteReader struct {}

func (r *InfiniteReader) Read(p []byte) (n int, err error) {
	for i:= range p {
		p[i] = 'A'
	}

	return len(p), nil
}

func readers() {
	// io.Reader is an interface that represents the ability to read data.
	// It has a single method Read that reads data into a byte slice.

	r := strings.NewReader("Hello, World!")

	b := make([]byte, 8) // Create a byte slice to hold the data
	for {
		// We read data from the reader into the byte slice
		n, err := r.Read(b)
		// n is the number of bytes read, err is an error if any occurred
		// Since our byte slice is 8 bytes long, we can read up to 8 bytes at a time.
		// If n < len(b), it means we read less than 8 bytes, which can happen if we reach the end of the data.


		// We will pass the first 8 bytes of the string to b, which are "Hello, W"
		// Then, we will read the next 8 bytes, which are "orld!" and then we will reach the end of the data.
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])

		// If err is io.EOF, it means we reached the end of the data.
		if err == io.EOF {
			break
		}
	}


	// Exercise Implementation

	// First, we will create a new type that implements the io.Reader interface.
	// Then, we will create a new value of that type.
	r2 := &InfiniteReader{}
	// We will read data from the reader into a byte slice.
	buf := make([]byte, 10)

	// In this case, since the reader is infinite,
	// n will be whatever the buffer can hold.
	n, err := r2.Read(buf)
	if err != nil && err != io.EOF {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Read %d bytes: %s\n", n, buf)
}