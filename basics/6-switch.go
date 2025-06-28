package main

import (
	"fmt"
	"runtime"
	"time"
)

func switches() {
	fmt.Print("Go runs on ")

	// Simple switch statement
	i := 2
	switch i {
		case 1:
			fmt.Println("one.")
		case 2:
			fmt.Println("two.")
		case 3:
			fmt.Println("three.")
		default: // Default is optional, but good practice
			fmt.Println("I don't know.")
	}

	// We can have multiple cases separated by commas
	switch time.Now().Weekday() {
		case time.Saturday, time.Sunday:
			fmt.Println("It's the weekend.")
		default:
			fmt.Println("It's a weekday.")
	}


	// Just like if statements or for loops, we can make a small statement before the switch block
	// in this case, "os := runtime.GOOS", which will only be available in the switch block
	switch os := runtime.GOOS; os {
		case "darwin":
			fmt.Println("macOS.")
		case "linux":
			fmt.Println("Linux.")
		default:
			fmt.Printf("%s.\n", os)
	}

	// We can also omit the condition
	// This serves as a cleaner way to write if-else statements
	t:= time.Now()
	switch {
		case t.Hour() < 12:  // if t.Hour() < 12
			fmt.Println("Good morning.")
		case t.Hour() < 16: // else if t.Hour() < 16
			fmt.Println("Good afternoon.")
		default: // else
			fmt.Println("Good evening.")
	}

	// We can also use a type switch to compare types instead of values
	whatAmI := func(i interface{}) {
		// t will be the type of i
		switch t := i.(type) {
			case bool:
				fmt.Println("I'm a bool")
			case int:
				fmt.Println("I'm an int")
			default:
				fmt.Printf("Don't know type %T\n", t)
		}
    }


	// If we want to break out of a switch statement, we can use the keyword "break"
	// If we want to break from a loop containing a switch statement,
	// We need to use a label for the loop and break out of that label
	Loop: // This is a label
		for i := range 10 {
			switch i {
				case 0:
					fmt.Println("zero")
					continue Loop // This will continue the loop named "Loop"
				case 1:
					fmt.Println("one")
				case 2:
					fmt.Println("two")
					break Loop // This will break out of the loop named "Loop"
				default:
					fmt.Println("other")
			}
		}
    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")
}