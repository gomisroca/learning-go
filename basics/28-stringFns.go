package main

import (
	"fmt"
	s "strings"
)

// The strings pkg contains many useful functions for working with strings

// alias fmt.Println as we will use it a lot
var p = fmt.Println

func stringFns() {
	p("Contains:          ", s.Contains("test", "es"))
	p("Count:             ", s.Count("test", "t"))
	p("Has Prefix:        ", s.HasPrefix("test", "te"))
    p("Has Suffix:        ", s.HasSuffix("test", "st"))
    p("Index:             ", s.Index("test", "e"))
    p("Join:              ", s.Join([]string{"a", "b"}, "-"))
    p("Repeat:            ", s.Repeat("a", 5))
    p("Replace All:       ", s.ReplaceAll("foo", "o", "0"))
    p("Replace N times:   ", s.Replace("foo", "o", "0", 1)) // Where n = 1
    p("Split:             ", s.Split("a-b-c-d-e", "-"))
    p("To Lower:          ", s.ToLower("TEST"))
    p("To Upper:          ", s.ToUpper("test"))
}