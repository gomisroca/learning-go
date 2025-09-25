package main

import (
	"crypto/sha256"
	"fmt"
)

// Usually used to compute short ids for binary or text data
func sha256Hashes() {
	s := "sha256 this string"

	// New hash
	h := sha256.New()
	// Write expected bytes (if we have a string, we use []byte(s) to get the bytes)
	h.Write([]byte(s))

	// Get the hash result as a byte slice
	// The arg is used to append to an existing slice
	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
} 