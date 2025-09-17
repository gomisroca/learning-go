package main

import (
	"fmt"
	"slices"
)

func sorting() {
	// Go's slices pkg has sorting for built-in and user-defined types

	// Sorting fns work for any ordered built-in type
	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	slices.Sort(ints)
	fmt.Println("Ints:", ints)

	// We can also check if a slice is already sorted
	s := slices.IsSorted(ints)
	fmt.Println("Sorted: ", s)
}