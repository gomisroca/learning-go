package main

import (
	b64 "encoding/base64" // rename the import
	"fmt"
)

func base64Encoding() {
	data := "abc123!?$*&()'-=@~"

	// Go support standard and URL-safe base64

	// Standard encoding
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)
	// Decoding may return an error
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	// URL-safe encoding
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}