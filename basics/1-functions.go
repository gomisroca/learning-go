package main

import "fmt"

func add(x, y int) int {
	return x + y
}

// A function can return multiple values
func vals(x, y int) (int, int) {
	return x + y, x - y
}

// Variadic functions can take any number of arguments
func sums(nums ...int) int {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}

	fmt.Println(total)
	return total
}

// Go also supports closures
func counter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// Go also supports recursion
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func functions() {
	fmt.Println("The sum of 3 and 4 is:", add(3, 4))
	fmt.Println("The sum of 10 and 20 is:", add(10, 20))
	fmt.Println("The sum of 100 and 200 is:", add(100, 200))

	x, y := vals(1, 2)
	fmt.Println(x, y)

	// If we don't use the variable, we can omit the variable name with _
	_, z := vals(1, 2)
	fmt.Println(z)

	// We can call the variadic function with any number of arguments
	sums(1, 2)
    sums(1, 2, 3)
	// And we can call it with a slice
    nums := []int{1, 2, 3, 4}
    sums(nums...)

	// We call counter and assign the result (a function) to a variable
	nextInt := counter()

	fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

	// Since the state of the function is local to the function, we can call it again and get a different result
    newInts := counter()
    fmt.Println(newInts())

	
    fmt.Println(fact(7)) // 5040

	// Anonymous functions can also be recursive, but we need to explicitly declare a var to store the function
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 { 
			return n
		}

		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7)) // 13
}