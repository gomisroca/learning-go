package main

import "fmt"

func arrays() {
	var a [3]int
	fmt.Println(a) // [0 0 0]
	a[0], a[1], a[2] = 1, 2, 3
	fmt.Println(a) // [1 2 3]

	// We can use a shorthand to declare an array
	i := [3]int{4, 5, 6}
	fmt.Println(i) // [4 5 6]

	// We can have the compiler infer the size of the array
	b := [...]int{4, 5, 6}
	fmt.Println(b) // [4 5 6]

	// We can also specify an index, the values in between will be zero'd
	b = [...]int{100, 2: 200}
	fmt.Println(b) // [100 0 200]

	var c [2]string
	c[0] = "Hello"
	c[1] = "World"
	fmt.Println(c) // [Hello World]
}