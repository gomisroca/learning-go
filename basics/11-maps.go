package main

import (
	"fmt"
	"maps"
)

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Google": {37.42202, -122.08408},
	"Facebook": {37.48484, -122.14838},
}

func keyExists(m map[string]Vertex, key string) {
	// We can check if a key exists in the map
	if v, ok := m[key]; ok {
		fmt.Printf("%s exists in the map: %v\n", key, v) // Google exists in the map: {32.45602 -122.08408}
	} else {
		fmt.Printf("%s does not exist in the map\n", key)
	}
}

func mapsInfo() {
	// A map maps keys to values

	// We can create an empty map with make
	n := make(map[string]int)

	// We can set key/value pairs like usual
	n["k1"] = 1
	n["k2"] = 2
	fmt.Println(n) // map[k1:1 k2:2]

	// Get the value for a key
	v1 := n["k1"]
	fmt.Println("v1:", v1) // 1

	// If the value does not exist, we get the zero value for the type
	v3 := n["k3"]
    fmt.Println("v3:", v3) // 0

	// We can also check if a key was present
	_, prs := n["k2"]
    fmt.Println("prs:", prs)

	// In this case, we are mapping the key "Bell Labs" to a Vertex of value {40.68433 -74.39967}
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m) // map[Bell Labs:{40.68433 -74.39967} Facebook:{37.48484 -122.14838} Google:{37.42202 -122.08408}]
	fmt.Println(m["Bell Labs"]) // {40.68433 -74.39967}

	m["Google"] = Vertex{32.45602, -122.08408} // We can update the value for a key
	fmt.Println(m["Google"]) // {32.45602 -122.08408}

	delete(m, "Facebook") // We can delete a key-value pair
	fmt.Println(m) // map[Bell Labs:{40.68433 -74.39967} Google:{32.45602 -122.08408}]

	// We can check if a key exists in the map
	keyExists(m, "Google") // Google exists in the map: {32.45602 -122.08408}
	keyExists(m, "Facebook") // Facebook does not exist in the map anymore

	// Finally, we can empty a map using clear
	clear(n)
	fmt.Println(n) // map[]

	// We can also use a shorthand to declare a map
	// The syntax is mapName := map[keyType]valueType{key1: value1, key2: value2, ...}
	o := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println(o) // map[a:1 b:2 c:3]

	// The maps package provides a lot of useful functions for maps
	p := map[string]int{"a": 1, "b": 2, "c": 3}
	if maps.Equal(p, o) {
		fmt.Println("p and o are equal")
	}
}