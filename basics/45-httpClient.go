package main

import (
	"bufio"
	"fmt"
	"net/http"
)

// Go has built-in support for HTTP clients and servers
func httpClient() {
	// Issue a GET request to a server
	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close() // Close the response body when done

	fmt.Println("Response status:", resp.Status) // Status should be "200 OK"

	// Print 5 lines of the response body
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}