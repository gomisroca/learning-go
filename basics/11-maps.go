package main

import "fmt"

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

func maps() {
	// A map maps keys to values

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
}