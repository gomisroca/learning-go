package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}


func (r *rot13Reader) Read(p []byte) (int, error) {
	// Create a buffer to read from the underlying reader
	buf := make([]byte, len(p))
	// Since r is the rot13Reader, we know it has an underlying io.Reader
	// Therefore, we can call Read on the underlying io.Reader
	n, err := r.r.Read(buf)

	// Then we apply ROT13 to each byte
	for i := range n {
		b := buf[i]
		switch {
		case 'A' <= b && b <= 'Z':
			p[i] = 'A' + (b-'A'+13)%26
		case 'a' <= b && b <= 'z':
			p[i] = 'a' + (b-'a'+13)%26
		default:
			p[i] = b
		}
	}

	return n, err
}

func rot13reader() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	// We pass s, which is a io.Reader, as the field r of rot13Reader
	r := rot13Reader{s}

	// io.Copy takes a writer and a reader, and it calls Read() on the reader repeatedly
	// In this case, we want to call Read() on the rot13Reader, which calls Read() on the underlying io.Reader
	// Finally, we will write the output to os.Stdout, which is the standard output
	io.Copy(os.Stdout, &r)
}