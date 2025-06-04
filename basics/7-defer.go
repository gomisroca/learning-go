package main

import (
	"fmt"
)

func defers() {
	// Defer statements are executed in LIFO order (Last In, First Out)
	// Meaning, the last defer statement will be executed first
	// They are executed after the surrounding function returns

	// This will be executed after defers() returns
	defer fmt.Println("This will be printed last")

	fmt.Println("This will be printed first")

	for i := range 3 {
		// This will be executed after the loop ends
		// Since it is LIFO, it will be printed in reverse order
		// So it will print 2, 1, 0
		defer fmt.Println(i)
	}
}