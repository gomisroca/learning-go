package main

import (
	"cmp"
	"fmt"
	"slices"
)

// Often we want to sort by something other than the default
// i.e sort strings by their length, instead of alphabetically
func sortingByFunctions() {
    fruits := []string{"peach", "banana", "kiwi"}

	// Implement a custom comparison function
	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	slices.SortFunc(fruits, lenCmp)
	fmt.Println("Fruits:", fruits)

	// We use the same method for non-built-in types

	type Person struct {
		name string
		age int
	}
	
    people := []Person{
        Person{name: "Jax", age: 37},
        Person{name: "TJ", age: 25},
        Person{name: "Alex", age: 72},
    }

	// Implement a custom comparison function
	// It will sort people by their age
	slices.SortFunc(people,
		func(a, b Person) int {
			return cmp.Compare(a.age, b.age)
		})
	fmt.Println("People:", people)
}