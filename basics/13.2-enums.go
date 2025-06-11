package main

import "fmt"

//go:generate stringer -type=ServerState

// Enums have a fixed number of possible values
type ServerState int // This is our enum type, which has an underlying type of int

const (
	StateIdle ServerState = iota // iota is a special constant that starts at 0 and increments by 1 for each enum value
	StateConnected
	StateError
	StateRetrying
)

// We can generate a map using the enum values as keys
var stateName = map[ServerState]string{
	StateIdle:     "Idle",
	StateConnected: "Connected",
	StateError:    "Error",
	StateRetrying: "Retrying",
}

// And then implement the Stringer interface:
func (s ServerState) String() string {
	if name, ok := stateName[s]; ok {
		return name // Would print "Idle", "Connected", "Error", or "Retrying"
	}
	// Fallback in case 's' is not a valid stateName key
	return fmt.Sprintf("UnknownState(%d)", s)
}
// In the real world, we would probably use go:generate to generate this automatically:
	// To do this, we add a comment to the top of the file, like this:
	// //go:generate stringer -type=ServerState
	// Then, we run "go generate" in the directory where the file is located.
	// This will generate a file called string.go, which contains the implementation of the Stringer interface.
	// Now we can use fmt.Println(StateIdle) or any other and it will print "StateIdle", "StateConnected", etc.

// A simple fn to emulate state transition in a server
func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state %s", s))
	}
}

func enums() {
	ns := transition(StateIdle)
	fmt.Println(ns) // Connected

	ns2 := transition(ns)
	fmt.Println(ns2) // Idle

	ns3 := transition(2) // We manually use 2 here, which corresponds to StateError due to the iota enum, we could have used transition(0) in the first case too, for example
	fmt.Println(ns3)
}