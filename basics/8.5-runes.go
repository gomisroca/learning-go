package main

import (
	"fmt"
	"unicode/utf8"
)

func runes() {
	// Go strings are UTF-8 encoded
	const s = "สวัสดี"

	// Since string are equivalent to []byte, len(s) will return the number of bytes used to store the string
	fmt.Println("Len: ", len(s)) // Len: 18

	for i := range len(s) {
		// We can access the bytes of a string using the index
		// For example here we get the hex value of the byte at index i
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	// Runes are equivalent to characters in other languages
	// To count how many runes are in a string, we can use the utf8 package
	fmt.Println("Rune count:", utf8.RuneCountInString((s))) // Rune count: 6

	// A range loop handles strings and will decode runes along the way
	for idx, runeValue := range s {
        fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}
	// The same can be achieved with utf8.DecodeRuneInString
	fmt.Println("\nUsing DecodeRuneInString")
    for i, w := 0, 0; i < len(s); i += w {
        runeValue, width := utf8.DecodeRuneInString(s[i:])
        fmt.Printf("%#U starts at %d\n", runeValue, i)
        w = width

        examineRune(runeValue)
    }

}

// Values enclosed in a single quote are rune literals
// We can compare a rune value to a rune literal directly
func examineRune(r rune) {
	  if r == 't' {
        fmt.Println("found tee")
    } else if r == 'ส' {
        fmt.Println("found so sua")
    }
}