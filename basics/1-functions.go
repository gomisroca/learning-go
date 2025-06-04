package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func functions() {
	fmt.Println("The sum of 3 and 4 is:", add(3, 4))
	fmt.Println("The sum of 10 and 20 is:", add(10, 20))
	fmt.Println("The sum of 100 and 200 is:", add(100, 200))
}