package main

import (
	"fmt"
	"math"
)

type Vtx struct {
	X, Y float64
}

type rect struct {
	width, height int
}

type myInt int

// We declare input before naming the method
// A method is a function with a receiver
// The receiver doesn't have to be a struct, it can be any type, we could say (v myInt) instead of (v Vtx), but the type has to be defined before the method (meaning types like int have to be defined as a new type, like myInt above)
// The receiver can also be a pointer to a value of the correct type
func (v Vtx) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// We can modify the receiver if we use a pointer receiver
// If we were to remove the *, we would be working with a copy of the value, not the original value
// This means that any changes made to the receiver would not be reflected in the original value
// For example in Abs(), we are using v as a copy, so we cannot modify it, but we use it to calculate and return another value
func (v *Vtx) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

// Since methods are functions, we could rewrite the above as:
// func Abs(v Vtx) float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }
// But this is not idiomatic Go, and we would lose the ability to call it like a method.
// We would have to call it like this: `Abs(v)` instead of `v.Abs()`.

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perimeter() int {
	return 2*(r.width + r.height)
}

// My own implementation of a method on a string-like type
type Word string

func (w *Word) Capitalize() {
	if len(*w) == 0 {
		return
	}
	runes := []rune(*w)
	if runes[0] >= 'a' && runes[0] <= 'z' {
		runes[0] = rune(runes[0] - ('a' - 'A'))
	}
	*w = Word(string(runes))
}

func methods() {
	v := Vtx{3, 4}
	fmt.Println(v.Abs()) // 5

	v.Scale(10)
	fmt.Println(v) // 30, 40

	w := Word("hello")
	fmt.Println(w) // hello
	w.Capitalize()
	fmt.Println(w) // Hello

	r := rect{10, 5}

	fmt.Println("Area: ", r.area()) // Area: 50
	fmt.Println("Perimeter: ", r.perimeter()) // Perimeter: 30

	// Go automatically handles conversion between values and pointers for method calls
	rp := &r
	fmt.Println("Area: ", rp.area()) // Area: 50
	fmt.Println("Perimeter: ", rp.perimeter()) // Perimeter: 30
}