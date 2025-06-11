package main

import "fmt"

// Go supports embedding, which means that a struct or interface can embed another struct or interface
// This is useful when we want to reuse some of the fields of a struct or interface, but we want to add some new fields
// For example, we could have a base struct that has a number field, and then we could embed it in a container struct that has a string field
type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("Base with num=%v", b.num)
}

type container struct {
	base // 👈 Here we are embedding the base struct
	str string
}

func structEmbedding() {
	// When creating structs with literals, we have to init the embeds explicitly
	co := container{
		base: base{ num: 1 }, // Require if base has no default values
		str: "some name",
	}
	// We can access the embedded fields using the dot operator directly from the container...
	fmt.Printf("co={num %v, str %v}\n", co.num, co.str)
	// ...or using the dot operator from the embedded struct
	fmt.Println("also num:", co.base.num)
	// Since container embeds base, the methods of base are also available on container
	fmt.Println("describe:", co.describe())

	// A powerful use of embeds is to implement interfaces:
	// We can create a describer interface that has a describe() method
	type describer interface {
		describe() string
	}
	// Since "co" already implements the describe() method via the embed of "base"...
	var d describer = co
	// ...we can use the method directly
	fmt.Println("describer:", d.describe())
}