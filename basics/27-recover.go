package main

import "fmt"

func mayPanic() {
	panic("a problem")
}

// We can recover from a panic by using recover()
func recovers()  {
	// recover must be called within a deferred fn
	// when the enclosing fn panics
	// the defer will activate and recover will catch the panic
	defer func() {
		// the return of recover is the error raised by the panic
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	// so now the enclosing fn will panic, which will trigger the defer and recover
	mayPanic()

	// this won't run, because the panic was caught by recover
	fmt.Println("After mayPanic()")
}
