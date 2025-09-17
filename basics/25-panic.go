package main

import "os"

func panics() {
	// A panic means something went unexpectedly wrong
	// or we aren't ready to handle it gracefully
	
	// panic("a problem") // Uncomment to see basic panic

	// A common use is to abort is a fn returns an error that we can't handle
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err) // we aren't handling this error, so we panic instead
	}
}