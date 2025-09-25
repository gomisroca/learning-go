package main

import (
	"fmt"
	"net"
	"net/url"
)

func urlParsing() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// Parse the URL
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	// Access the scheme
	fmt.Println(u.Scheme)
	// Access user data
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p ,_ := u.User.Password()
	fmt.Println(p)

	// Access host data
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	// Access the path and fragment
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	// Access query params in k=v format
	fmt.Println(u.RawQuery)
	// Map the query params
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}