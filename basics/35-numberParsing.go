package main

import (
	"fmt"
	"strconv"
)

// We can parse numbres from strings
func numberParsing() {
	// Parse floats with arg2 bits of precision
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// Parse ints with arg2 base and arg3 bits of precision
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)
	// Can use hex-format numbers
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// Parse Uints
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// Parse atoi, which is for basic base-10 parsing
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	// Parse fns error on bad input
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}