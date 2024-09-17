// String Functions
//
// The standard library's strings package provides many useful string-related
// functions. Here are some examples to give you a sense of the package.
//
// Run: go run ./54_string_functions
package main

import (
	"fmt"
	s "strings"
)

// We alias fmt.Println to a shorter name as we'll use it a lot below.
var p = fmt.Println

func main() {
	// Here's a sample of the functions available in strings. Since these are
	// functions from the package, not methods on the string object itself, we need
	// to pass the string in question as the first argument to the function.
	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))

	// Trimming and splitting on whitespace.
	p("TrimSpace: ", s.TrimSpace("  padded  "))
	p("Trim:      ", s.Trim("--x--", "-"))
	p("TrimPrefix:", s.TrimPrefix("v1.2", "v"))
	p("Fields:    ", s.Fields("  a  b c "))

	// Cut splits around the first separator and reports whether it was found.
	key, val, ok := s.Cut("name=gopher", "=")
	p("Cut:       ", key, val, ok)

	// EqualFold compares case-insensitively.
	p("EqualFold: ", s.EqualFold("Go", "GO"))

	// strings.Builder builds a string efficiently, without copying on every +.
	var b s.Builder
	for i := range 3 {
		fmt.Fprintf(&b, "%d,", i)
	}
	b.WriteString("done")
	p("Builder:   ", b.String())

	// A Replacer performs several replacements in one pass.
	r := s.NewReplacer("<", "&lt;", ">", "&gt;")
	p("Replacer:  ", r.Replace("<b>hi</b>"))
}
