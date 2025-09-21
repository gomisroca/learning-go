package main

import (
	"bytes"
	"fmt"
	"regexp"
)

// Go has built-in support for regular expressions
func regex() {
	// Check if a string matches a regular expression
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
    fmt.Println(match)

	// Normally we will need to compile an optimized regexp struct
	r, _ := regexp.Compile("p([a-z]+)ch")

	// These structs have many methods:
	// Match a string
	fmt.Println(r.MatchString("peach"))
	// Find a match
	fmt.Println(r.FindString("peach punch"))
	// Find first match, but return the start and end idx of the match
	fmt.Println("idx:", r.FindStringIndex("peach punch"))
	// Find a submatch (in this case, it would also look for matches for [a-z]+)
	fmt.Println(r.FindStringSubmatch("peach punch"))
	// Find idx of submatch
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// We can add All to these functions to find all matches
	fmt.Println(r.FindAllString("peach punch pinch", -1)) // Second arg limits the number of matches

	// We can also match non-string types
	fmt.Println(r.Match([]byte("peach")))

	// We can use MustCompile instead of Compile, which will panic if there is an error
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r.MatchString("peach"))

	// We can replace matches
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// And we can use Func to replace matches with a function
	in := []byte("a peach")
    out := r.ReplaceAllFunc(in, bytes.ToUpper) // matches will be uppercased
    fmt.Println(string(out))
}