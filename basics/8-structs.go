package main

import "fmt"

type person struct {
	name string
	age  int
}

// This function creates a new person struct with the given name and age
func newPerson(name string, age int) *person {
	p := person{name, age}
	return &p  // We return a pointer to the struct, so we can modify it later
}

func structs() {
	fmt.Println(newPerson("Jon", 30)) // {Jon 30}

	// We can also shorthand the creation of a struct
	fmt.Println(person{"Alice", 30}) // {Alice 30}

	// We can create a struct and point to its pointer
	fmt.Println(&person{"Tom", 23}) // &{Tom 23}

	// They can be stored in variables
	p1 := person{"Bob", 25}
	p2 := person{"Charlie", 30}


	// We can access their fields using the dot notation
	fmt.Println(p1.name) // Bob
	fmt.Println(p2.age)  // 30

	// We can also create a pointer to a struct
	p := &p1
	// And access its fields using the pointer, which modifies the original struct
	p.name = "Bob the Builder"
	fmt.Println(p1.name) // Bob the Builder

	// We can also create pointers to structs
	// And we can create structs without specifying the fields, in which case the fields are initialized to their zero values
	p3pointer := &person{name: "Dave"}
	fmt.Println(p3pointer.name) // Dave


	// If a struct is only used for a single value, we don't have to give it a name
	dog := struct {
		name string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog) // {Rex true}
}