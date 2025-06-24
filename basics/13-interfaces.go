package main

import (
	"fmt"
	"math"
	"reflect"
)

// An interface is a type that defines a set of methods
// A type implements an interface by implementing its methods

// 🥸 “I don’t care what you are, but if you have my methods — then you're a geometry.”
// So any type that implements area() and perimeter() is a 'geometry', regardless of anything else.
type geometry interface { 
	area() float64
	perimeter() float64
	// sin() float64 // If we add another method, it will have to be implemented by the structs that want to implement the 'geometry' interface
}

type triangle struct {
	a, b, c float64
}

type circle struct {
	radius float64
}

// To implement an interface, we need to implement all the methods on the struct
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
    return 2 * math.Pi * c.radius
}
func (c circle) String() string {
	return fmt.Sprintf("Circle with radius %g", c.radius)
}

func (t triangle) area() float64 {
    return t.a * t.b * t.c / 2
}
func (t triangle) perimeter() float64 {
    return t.a + t.b + t.c
}
func (t triangle) String() string {
	return fmt.Sprintf("Triangle with sides %g, %g, %g", t.a, t.b, t.c)
}

// We will be able to pass any type that implements the geometry interface to this function
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("Area:", g.area())
	fmt.Println("Perimeter:", g.perimeter())
}

// We can use type assertions to check the type
func detectCircle(g geometry) {
    if c, ok := g.(circle); ok {
        fmt.Println("This is a circle with a radius of", c.radius)
    } else {
		// Here, %T would print the type of reflect.TypeOf(g).Name(), which would be a string
		// %v would print the value of reflect.TypeOf(g).Name(), which would be whatever the type of g is
		fmt.Printf("This is not a circle, this is a %v!\n", reflect.TypeOf(g).Name())
	}

	// We can do the same with a type switch
	// switch t := g.(type) {
	// case circle:
	// 	fmt.Println("This is a circle with a radius of", t.radius)
	// default:
	// 	fmt.Printf("This is not a circle, this is a %v!\n", reflect.TypeOf(t).Name())
	// }
}

func interfaces() {
	t := triangle{3, 4, 5}
	c := circle{5}

	// If the values of t or c were not set, they would be nil
	// If we tried calling t.area() and it was nil, it would panic
	// We usually check inside the method if the receiver is nil, and handle it gracefully

	measure(t)
	measure(c)

	detectCircle(t)
	detectCircle(c)

	// We can also check if an interface holds a specific type using a type assertion
	var i any = "hello" 
		// empty interface can hold any type
	 	// any is an alias for interface{} (empty interface) and is equivalent to interface{} in all ways.
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