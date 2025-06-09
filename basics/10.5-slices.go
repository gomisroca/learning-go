package main

import "fmt"

func slices() {
	// Slices are similar to arrays, but they are mutable
	// An uninitialized slice is nil, and its length and capacity are both zero
	var s []string
    fmt.Println("uninit:", s, s == nil, len(s) == 0) // uninit: [] true true

	// We can create a slice with a length and capacity using make
	s = make([]string, 3)
    fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s)) // emp: [  ] len: 3 cap: 3

	// We can declare and initialize a slice at the same time
 	t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)

	// We can get and set like with arrays
	s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s) // set: [ a b c ]
    fmt.Println("get:", s[2]) // get: c
	fmt.Println("len:", len(s)) // len: 3

	// We can also append
	s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("apd:", s) // apd: [ a b c d e f ]

	// We can also copy the slice
	c := make([]string, len(s)) // Here we are making an empty slice of the same length as s
    copy(c, s) // And copying the contents of s into c
    fmt.Println("cpy:", c) // cpy: [ a b c d e f ]

	// We can also get a slice of a slice
	l := s[2:5] // the low bound is included, the high bound is not
    fmt.Println("sl1:", l) // sl1: [ c d e ]

	// We can loop over a slice using range
	for i, v := range c {
		fmt.Printf("c[%d] = %s\n", i, v) // c[0] = a, c[1] = b, ...
	}

	// We can omit the index or value if we don't need it
	for i := range c {
		// This essentially prints the same as above
		fmt.Printf("c[%d] = %s\n", i, c[i]) // c[0] = a, c[1] = b, ...
	}
	for _, v := range c {
		fmt.Printf("%s\n", v) // a b c d e f
	}

	// We can create multi-dimensional data structures with slices
	twoD := make([][]int, 3) // We can make a slice of slices as the outer slice
    for i := range 3 {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen) // Then, another slice as the inner slice
        for j := range innerLen {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD) // 2d:  [[0] [1 2] [2 3 4]]
}