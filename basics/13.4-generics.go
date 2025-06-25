package main

import "fmt"

// Generics are a way to define a type that can be used with any other type
// They are defined using the keyword type, followed by the name of the type, followed by the name of the type parameter, and then the type definition

// For example, we can define a generic list type that can hold any type
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct { 
	next *element[T]
	val T
}

// We can define methods on the generic types just like we would on any other type
func (list *List[T]) Push(val T) {
	// If there is no tail, it means that the list is empty
	if list.tail == nil {
		// We create a new element and set its pointer as the head and tail
		list.head = &element[T]{val: val}
		list.tail = list.head
	} else {
		// Else, we create a new element and set its pointer as the next of the tail
		list.tail.next = &element[T]{val: val}
		// And move the tail
		list.tail = list.tail.next
	}
}

func (lst *List[T]) AllElements() []T {
	// Init a slice of the same type as the elements in the list
    var elems []T
	// Start from the head and go to next until we reach the end, appending each element to the slice
    for e := lst.head; e != nil; e = e.next {
        elems = append(elems, e.val)
    }
    return elems
}

func generics() {
	lst := List[int]{}
    lst.Push(10)
    lst.Push(13)
    lst.Push(23)
    fmt.Println("list:", lst.AllElements())
}