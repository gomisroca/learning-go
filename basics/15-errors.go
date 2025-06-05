package main

import (
	"fmt"
	"time"
)

// error is a built-in interface like fmt.Stringer
// type error interface {
//     Error() string
// }
// Like Stringer, fmt will

type MyError struct {
	Message string
	Time    time.Time
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Error: %s at %s", e.Message, e.Time)
}

func division(x, y int) (int, error) {
	if y == 0 {
		return 0, &MyError{
			Message: "Division by zero is not allowed",
			Time:    time.Now(),
		}
	}
	return x / y, nil
}

func errors() {
	res1, err := division(10, 2) // This will not return an error
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res1)
	}

	res2, err := division(10, 0) // This will return an error
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res2)
	}
}