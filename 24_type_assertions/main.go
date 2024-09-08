// Type Assertions
//
// An interface value holds a concrete value and its type. A type assertion gets the
// concrete value back out, and a type switch branches on which type is inside.
//
// Run: go run ./24_type_assertions
package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	var i any = "hello"

	// x.(T) asserts that i holds a T and returns the value.
	s := i.(string)
	fmt.Println(s)

	// The comma-ok form reports failure instead of panicking.
	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// Without ok, a failed assertion panics.
	func() {
		defer func() { fmt.Println("recovered:", recover()) }()
		_ = i.(float64)
	}()

	// A type switch tests several types in order. In each case, v has that case's
	// type.
	for _, x := range []any{42, "go", 3.5, nil, []int{1}} {
		switch v := x.(type) {
		case int:
			fmt.Println("int, doubled:", v*2)
		case string:
			fmt.Println("string of length", len(v))
		case nil:
			fmt.Println("nil")
		default:
			fmt.Printf("something else: %T\n", v)
		}
	}

	// Asserting to an interface type checks for behavior. This is how code finds
	// out whether a value has an optional method, such as Close.
	var r io.Reader = strings.NewReader("data")
	if _, ok := r.(io.Closer); !ok {
		fmt.Println("strings.Reader has no Close method")
	}
	if wt, ok := r.(io.WriterTo); ok {
		fmt.Printf("but it is a WriterTo: %T\n", wt)
	}
}
