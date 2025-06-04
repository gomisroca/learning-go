package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func structs() {
	// We can define a struct to group related data together
	fmt.Println(Person{"Alice", 30}) // {Alice 30}

	// They can be stored in variables
	p1 := Person{"Bob", 25}
	p2 := Person{"Charlie", 30}


	// We can access their fields using the dot notation
	fmt.Println(p1.Name) // Bob
	fmt.Println(p2.Age)  // 30

	// We can also create a pointer to a struct
	p := &p1
	// And access its fields using the pointer, which modifies the original struct
	p.Name = "Bob the Builder"
	fmt.Println(p1.Name) // Bob the Builder

	// We can also create pointers to structs
	// And we can create structs without specifying the fields, in which case the fields are initialized to their zero values
	p3pointer := &Person{Name: "Dave"}
	fmt.Println(p3pointer.Name) // Dave

}