package main

import (
	"fmt"
	"math"
)

func ifElse() {
	x := 10

	// Like For Loops, no ( ) around the condition, but { } are mandatory
	if x > 5 {
		fmt.Println("x is greater than 5")
	} else if x < 5 {
		fmt.Println("x is less than 5")
	} else {
		fmt.Println("x is equal to 5")
	}


	// We can also execute a short statement before the if:
	// Vars declared here are only available in the if block
	if v := math.Pow(2, 3); v > 10 {
		fmt.Println("2^3 is greater than 10")
	} else {
		fmt.Println("2^3 is not greater than 10")
	}
	// return v // This would cause an error, v is not in scope here

	// Logical operators are also available
	if x > 5 && x < 10 {
		fmt.Println("x is between 5 and 10")
	}
	if x == 5 || x == 10 {
		fmt.Println("x is either 5 or 10")
	}

	// There is no ternary operator in Go (ex. x > 5 ? true : false) We need to use a full if statement
}