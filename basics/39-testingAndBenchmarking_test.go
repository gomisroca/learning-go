package main

import (
	"fmt"
	"testing"
)

// The testing package provides a set of functions for writing tests
// We can run test with go test

// Simple test
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// A test fn must start with Test
func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		t.Errorf("IntMin(2, -2) = %d; want -2", ans) // Test failure
	}
}

// It's idiomatic to use table-driven tests, 
// where test inputs/outputs are in a table
// we then loop over the table
func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
        a, b int
        want int
    }{
        {0, 1, 0},
        {1, 0, 0},
        {2, -2, -2},
        {0, -1, -1},
        {-1, 0, -1},
    }

	// t.Run runs "subtests" for each entry in the table
    // We can see the subtests with go test -v
	for _, tt := range tests {
        testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
        t.Run(testname, func(t *testing.T) {
            ans := IntMin(tt.a, tt.b)
            if ans != tt.want {
                t.Errorf("got %d, want %d", ans, tt.want)
            }
        })
    }
}

// Benchmark tests are used to measure performance
// It will automatically run the test fn many times
func BenchmarkIntMin(b *testing.B) {
    for b.Loop() {
        IntMin(1, 2)
    }
}