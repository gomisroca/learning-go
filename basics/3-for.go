package main

import "fmt"

func forLoop() {
	sum := 0

	// Pretty much the same as JS, PY, etc.
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println(sum)
	}

	// We can also use range to accomplish the "do this N times" loop
  	for i := range 3 {
        fmt.Println("range", i)
    }

	// If we don't need an init or step value, we can omit them
	// This creates a while loop
	for sum < 100 {
		sum += sum
		fmt.Println(sum)
	}

	// Can also use continue and break
	for i := range 10 {
		if i % 2 == 0 {
			continue
		}
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
		
	// We can also create an infinite loop by omitting the condition
	for {
		fmt.Println("I'm an infinite loop")
		break; // Remove this line to make it infinite
	}
}