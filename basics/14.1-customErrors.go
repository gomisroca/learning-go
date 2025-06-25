package main

import (
	"errors"
	"fmt"
)

// We can define custom error types by implementing the Error() method.
// This allows us to provide additional context or information about the error.

type argError struct {
	arg  int
	msg  string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.msg)
}

func f(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with 42"}
	}
	return arg + 3, nil
}

func customErrors() {
 	_, err := f(42)

    var ae *argError

	// errors.As() is used to check if an error is of a specific type or matches a specific error value.
    if errors.As(err, &ae) {
        fmt.Println(ae.arg) // 42
        fmt.Println(ae.msg) // can't work with 42
		fmt.Println(err)    // 42 - can't work with 42
    } else {
        fmt.Println("err doesn't match argError")
    }
}