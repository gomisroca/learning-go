// C flashbacks

package main

import "fmt"

func pointers() {
	// Pointers are variables that store the memory address of another variable
	// They are useful for sharing data between functions without copying it

	i, j := 42, 2701

	// We can declare pointers in verbose or non-verbose ways
	var k *int = &j
	p := &i

	// We can access the address directly
	fmt.Println(p) // &{2701}
	fmt.Println(k) // &{42}
	// We don't need to assign the address to a variable to access it
	fmt.Println(&i) // &{42} 

	// With * we can access the value stored at the address
	fmt.Println(*p) // 42
	fmt.Println(*k) // 2701

	// We can also change the value stored at the address
	*p, *k = 100, 1000
	fmt.Println(*p) // 100
	fmt.Println(*k) // 1000

	h := 1
	fmt.Println("before:", h)
	l := &h
	*l = 2
	fmt.Println("after:", h)
}