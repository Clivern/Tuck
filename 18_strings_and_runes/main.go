// Strings and Runes
//
// A Go string is a read-only slice of bytes. The language and the standard library
// treat strings as containers of text encoded in UTF-8. In other languages strings
// are made of characters; in Go, a character is called a rune, an integer that
// represents a Unicode code point.
//
// Run: go run ./18_strings_and_runes
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// s is a string holding the word "hello" in Thai. String literals are UTF-8
	// encoded text.
	const s = "สวัสดี"

	// Since strings are bytes, len gives the number of raw bytes.
	fmt.Println("Len:", len(s))

	// Indexing a string gives the byte at each index.
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	// To count runes, use the utf8 package. Some Thai characters are made of
	// several code points, so the count may surprise you.
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// range decodes each rune and its byte offset.
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	// The same iteration with utf8.DecodeRuneInString.
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		// A rune can be passed to a function and compared to a rune literal.
		examineRune(runeValue)
	}

	// Strings cannot be changed in place. Convert to []rune or []byte, change that,
	// and convert back.
	r := []rune("hello")
	r[0] = 'j'
	fmt.Println(string(r))

	// Raw string literals use backquotes: no escapes, and they can span lines.
	fmt.Println(`C:\no\escapes`)
}

func examineRune(r rune) {
	// Values in single quotes are rune literals.
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
