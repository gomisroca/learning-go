package main

import (
	"errors"
	"fmt"
	"time"
)

// error is a built-in interface like fmt.Stringer
// type error interface {
//     Error() string
// }
// Like Stringer, fmt will

type DivisionError struct {
	Message string
	Time    time.Time
}

func (err *DivisionError) Error() string {
	return fmt.Sprintf("customError: %s at %s", err.Message, err.Time.Format(time.RFC3339))
}

func division(x, y int) (int, error) {
	if y == 0 {
		return 0, &DivisionError{
			Message: "Division by zero is not allowed",
			Time:    time.Now(),
		}
	}
	return x / y, nil
}

// We don't need to define a custom error type if we don't need additional context.
// We can use the built-in errors.New() function to create a simple error.
// This is useful for simple error messages without additional context.
// The error returned by errors.New() is of type error, which implements the Error() method.
func sum(x, y int) (int, error) {
	if x < 0 || y < 0 {
		return 0, errors.New("errors.New(): sum cannot be calculated for negative numbers")
	}
	return x + y, nil
}

// Furthermore, we can use fmt.Errorf to create formatted error messages.
// This allows us to include dynamic values in the error message.
// It is a good middle ground between custom error types and simple error messages.
func fmtError(x, y int) error {
	return fmt.Errorf("fmt.Errorf: cannot calculate sum of %d and %d: negative numbers are not allowed", x, y)
}

func errorHandling() {
	res1, err := division(10, 0) // This will return a custom error
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res1)
	}

	res2, err := sum(10, -10) // This will return a regular errors.New() error
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res2)
	}

	err = fmtError(10, 20) // This will return a formatted fmt.Errorf() error
	if err != nil {
		fmt.Println(err)
	}
}