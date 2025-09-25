package main

import (
	"fmt"
	"math/rand/v2"
)

// math/rand/v2 provides pseudo-random number generation
func randomNumbers() {
	// rand.Intn(100) returns a random integer between 0 and 100
	fmt.Print(rand.IntN(100), ",") // comma-separated
	fmt.Print(rand.IntN(100))
	fmt.Println()

	// rand.Float64() returns a random float between 0.0 and 1.0
	fmt.Println(rand.Float64())
	fmt.Println((rand.Float64() * 5) + 5, ",")
	fmt.Println((rand.Float64() * 5) + 5)
	fmt.Println()

	// We can use a known seed
	s2 := rand.NewPCG(42, 1024)
	r2 := rand.New(s2)
	fmt.Print(r2.IntN(100), ",")
	fmt.Print(r2.IntN(100))
	fmt.Println()

	// Since we are using the same seed, we get the same sequence of numbers
    s3 := rand.NewPCG(42, 1024)
    r3 := rand.New(s3)
    fmt.Print(r3.IntN(100), ",")
    fmt.Print(r3.IntN(100))
    fmt.Println()
}