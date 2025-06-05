package main

import "fmt"

// An interface is a type that defines a set of methods
// A type implements an interface by implementing its methods
type Abser interface {
	Abs() float64
}

func interfaces() {
	var a Abser // a is an interface that can hold any type that implements the Abs() method

	// If the valus of a is not set, it will be nil
	// If we try to call a.Abs() now, it will panic because a is nil
	// We usually check inside the method if the receiver is nil, and handle it gracefully

	// i := myInt(5)
	v := Vtx{3, 4}

	// In lesson 12, we defined a method Abs() on Vtx, so we can assign a Vtx to an Abser

	// a = i // This will not work, because myInt does not implement the Abs() method
	a = v // This works because Vtx implements the Abs() method

	fmt.Println(a.Abs()) // 5

	// We can also check if an interface holds a specific type using a type assertion
	var i interface{} = "hello" // empty interface can hold any type
	s := i.(string) // here we are assuming that i holds a string
	fmt.Println(s)
	// If we are not sure that i holds a string, we can use a comma-ok idiom
	s, ok := i.(string)
	fmt.Println(s, ok)

	// If ok is false, it means that i does not hold the type, for example, "hello" is not a float64
	f, ok := i.(float64)
	fmt.Println(f, ok)
	// and if we try to use f, it will panic

	// Easier done with a type switch
	switch v := i.(type) {
	case string:
		fmt.Println("string:", v)
	case int:
		fmt.Println("int:", v)
	default:
		fmt.Println("unknown type:", v)
	}
}