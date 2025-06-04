package main

import "fmt"

func vars() {
	// Multiple variables can be declared at once
	var i, j int = 1, 2
	// If the type is clear from the context, we can omit it
	var h, g = "hello", "world"

	// Less verbose way to declare variables
	k := 3
	
	fmt.Println("i =", i, "j =", j, "k =", k, "h =", h, "g =", g)	
}