package main

import (
	"fmt"

	"github.com/google/uuid"
)

func packages() {
	id := uuid.New()
	fmt.Println("Generated UUID:", id.String()) // Could also do uuid.New().String() directly
	
	// If we add a package and want to remove it,
	// we can use `go mod tidy` to clean up unused packages
}