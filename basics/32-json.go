package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Go has built-in support for JSON encoding and decoding, and custom data types

// Types to demonstrate encoding and decoding
type response1 struct {
	Page int
	Fruits []string
}
// Only exported fields will be encoded/decoded in JSON. They must be capitalized
type response2 struct {
	Page int `json:"page"`
	Fruits []string `json:"fruits"`
}


func jsons() {
	// Encoding basic types
	// bool
	bolB, _ := json.Marshal(true)
    fmt.Println(string(bolB))
	// int
	intB, _ := json.Marshal(1)
    fmt.Println(string(intB))
	// float64
    fltB, _ := json.Marshal(2.34)
    fmt.Println(string(fltB))
	// string
    strB, _ := json.Marshal("gopher")
    fmt.Println(string(strB))

	// Slices and maps, which are encoded as JSON arrays and objects
	slcD := []string{"apple", "peach", "pear"}
    slcB, _ := json.Marshal(slcD)
    fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
    mapB, _ := json.Marshal(mapD)
    fmt.Println(string(mapB))

	// JSON can encode custom types, it will use exported fields as the JSON keys
	res1D := &response1{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res1B, _ := json.Marshal(res1D)
    fmt.Println(string(res1B))

	// In the declaration of response2 we added json tags, 
	// which will be used as the JSON keys instead of the exported fields names
	res2D := &response2{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res2B, _ := json.Marshal(res2D)
    fmt.Println(string(res2B))


	// Decoding
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	// We need a variable where we can store the decoded data
	var dat map[string]any

	// Decode while checking for errors
	if err := json.Unmarshal(byt, &dat); err != nil {
        panic(err)
    }
    fmt.Println(dat)

	// To use the data, we need to convert it to the correct type
	num := dat["num"].(float64)
    fmt.Println(num)
	// Accessing nested data requires a series of conversions
	strs := dat["strs"].([]any)
    str1 := strs[0].(string)
    fmt.Println(str1)

	// Can also decode into a struct
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
    res := response2{}
    json.Unmarshal([]byte(str), &res)
    fmt.Println(res)
    fmt.Println(res.Fruits[0])

	// Can encode/decode directly to os.Stdout or HTTP body responses
	enc := json.NewEncoder(os.Stdout)
    d := map[string]int{"apple": 5, "lettuce": 7}
    enc.Encode(d)

	dec := json.NewDecoder(strings.NewReader(str))
    res1 := response2{}
    dec.Decode(&res1)
    fmt.Println(res1)
	
	// A real world example from the gasthaus backend:
		// type loginRequest struct {
		// 	Email    string `json:"email"`
		// 	Password string `json:"password"`
		// }

		// type loginResponse struct {
		// 	Token string `json:"token"`
		// }

		// func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
		// 	var req loginRequest
		//  > Decode the request body into req
		// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		// 		http.Error(w, "Invalid request body", http.StatusBadRequest)
		// 		return
		// 	}
		//
		// > Do some auth stuff
		//
		// > Encode a response of type loginResponse and write it to the response body
		// json.NewEncoder(w).Encode(loginResponse{Token: tokenString})
}