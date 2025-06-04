package main

import "fmt"

func forLoop() {
	sum := 0

	// Pretty much the same as JS, PY, etc.
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println(sum)
	}

	// If we don't need an init or step value, we can omit them
	// This creates a while loop
	for sum < 100 {
		sum += sum
		fmt.Println(sum)
	}

	// We can also create an infinite loop by omitting the condition
	for {
		fmt.Println("I'm an infinite loop")
	}
}