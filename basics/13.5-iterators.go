package main

import (
	"fmt"
	"iter"
	"slices"
)

// We will use the List, element and Push methods from the generics example

func (lst *List[T]) All() iter.Seq[T] {
	// The iterator function takes another func as an argument, which is called a yield function
	// The yield function is called for each element in the list
	// If the yield function returns false, the iteration stops
	return func(yield func(T) bool) {
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

// Iteration doesn't require an underlying data structure
// Doesn't have to be finite either
func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1

		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func iterators() {
    lst := List[int]{}
    lst.Push(10)
    lst.Push(13)
    lst.Push(23)

	// list.All() returns an iterator function that can be used to iterate over the elements of the list
    for e := range lst.All() {
        fmt.Println(e) //  10, 13, 23
    }

	// Slices has many useful functions that can be used with iterators
	// Collect collects all elements from the iterator into a slice
    all := slices.Collect(lst.All())
    fmt.Println("all:", all) // [10 13 23]

	// Once the loop hits a break or early return, the yield function of the iterator stops being called
    for n := range genFib() {

        if n >= 10 {
            break
        }
        fmt.Println(n) // 1, 1, 2, 3, 5, 8
    }
}