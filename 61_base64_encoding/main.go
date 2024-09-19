// Base64 Encoding
//
// Go provides built-in support for base64 encoding and decoding, and for hex encoding in
// encoding/hex.
//
// Run: go run ./61_base64_encoding
package main

import (
	b64 "encoding/base64"
	"encoding/hex"
	"fmt"
)

// This syntax imports the encoding/base64 package with the b64 name instead of the
// default base64. It'll save us some space below.
func main() {
	// Here's the string we'll encode/decode.
	data := "abc123!?$*&()'-=@~"

	// Go supports both standard and URL-compatible base64. Here's how to encode using
	// the standard encoder. The encoder requires a []byte so we convert our string
	// to that type.
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	// Decoding may return an error, which you can check if you don't already know
	// the input to be well-formed.
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	// This encodes/decodes using a URL-compatible base64 format.
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))

	// The string encodes to slightly different values with the standard and URL
	// base64 encoders (trailing + vs -) but they both decode to the original string
	// as desired. RawURLEncoding also drops the = padding, as JWTs do.
	fmt.Println(b64.RawURLEncoding.EncodeToString([]byte("hi")))

	// Invalid input is reported as an error.
	_, err := b64.StdEncoding.DecodeString("not base64!")
	fmt.Println("error:", err)

	// Hex encoding writes two characters per byte.
	h := hex.EncodeToString([]byte("Go"))
	back, _ := hex.DecodeString(h)
	fmt.Println(h, string(back))
}
