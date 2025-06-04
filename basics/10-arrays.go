package main

import "fmt"

func arrays() {
	var a [3]int
	a[0], a[1], a[2] = 1, 2, 3
	fmt.Println(a) // [1 2 3]

	b := [...]int{4, 5, 6}
	fmt.Println(b) // [4 5 6]

	var c [2]string
	c[0] = "Hello"
	c[1] = "World"
	fmt.Println(c) // [Hello World]

	d := [6]int{1, 2, 3, 4, 5, 6}
	e := d[1:4]
	fmt.Println(e) // [2 3 4]

	// We can modify the original array through the slice
	e[0] = 10
	fmt.Println(d) // [1 10 3 4 5 6]

	// We can create dynamically-sized arrays using make
	f := make([]int, 5) // Creates a slice of length 5 with zero values
	fmt.Printf("%s len=%d cap=%d %v\n", "f", len(f), cap(f), f) // len=5 cap=5 [0 0 0 0 0]

	g := make([]int, 0, 10) // Creates a slice of length 5 with capacity 10
	fmt.Printf("%s len=%d cap=%d %v\n", "g", len(g), cap(g), g) // len=0 cap=10 []

	h := g[:2] // Slicing the slice
	fmt.Printf("%s len=%d cap=%d %v\n", "h", len(h), cap(h), h) // len=2 cap=10 [0 0]

	// We can append to a slice
	h = append(h, 1, 2, 3)
	fmt.Printf("%s len=%d cap=%d %v\n", "h", len(h), cap(h), h) // len=5 cap=10 [0 0 1 2 3]

	// We can loop over a slice using range
	for i, v := range c {
		fmt.Printf("c[%d] = %s\n", i, v) // c[0] = Hello, c[1] = World
	}

	// We can omit the index or value if we don't need it
	for i:= range c {
		// This essentially prints the same as above
		fmt.Printf("c[%d] = %s\n", i, c[i]) // c[0] = Hello, c[1] = World
	}
	for _, v := range c {
		fmt.Printf("%s ", v) // Hello World
	}
}