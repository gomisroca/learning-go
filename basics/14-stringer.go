package main

import "fmt"

type IPAddr [4]byte

// In GO, any type that implments a String() string method, will make fmt print the return value of that method, instead of the default representation of the type.
// This is called Stringer interface, and it is defined in the fmt package
// The Stringer interface is defined as:
// type Stringer interface {
//     String() string
// }

// In the background, fmt uses type assertion to check if the type implements the Stringer interface
// Essentially, does the value satisfy the Stringer interface, which means does it have a String() method that returns a string.
// fmt.Stringer is simply asking, "does this type implement the String() method?"
//
// if stringer, ok := value.(fmt.Stringer); ok {
//     return stringer.String()
// }


// Here we implement the String() method for the IPAddr type, so that when we print an IPAddr, it will print the IP address in a human-readable format.
func (ip IPAddr) String() string {
    return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

// My own implementation of a Stringer interface
type User struct { 
	Name string
	Age  int
}

func (u User) String() string {
	// This method implements the Stringer interface, so fmt will use this method to print Hello
	return fmt.Sprintf("Hello, %s, you are %d years old!", u.Name, u.Age)
}


func stringer() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip) // This will call the String() method of IPAddr, because fmt knows that IPAddr implements the Stringer interface
		// Output:
		// loopback: 127.0.0.1 
		// googleDNS: 8.8.8.8

		// If we didn't implement the String() method, we would get:
		// loopback: 127 0 0 1
		// googleDNS: 8 8 8 8
		// This is because fmt would use the default representation of the IPAddr type, which is an array of bytes.
	}

	
	john := User{"John", 42}
	fmt.Println(john) // Hello, John, you are 42 years old!
	// This will call the String() method of User, because fmt knows that User implements the Stringer interface
}


