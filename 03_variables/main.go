// Variables
//
// Variables are declared explicitly and the compiler checks their types, for example
// that a function receives arguments of the right kind.
//
// Run: go run ./03_variables
package main

import "fmt"

func main() {
	// var declares one or more variables.
	var a = "initial"
	fmt.Println(a)

	// You can declare several at once.
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go infers the type of an initialized variable.
	var d = true
	fmt.Println(d)

	// A variable declared without a value gets its zero value: 0 for numbers, "" for
	// strings, false for booleans, and nil for pointers, slices, and maps.
	var e int
	var f string
	var g []int
	fmt.Println(e, f == "", g == nil)

	// Inside a function, := is shorthand for declaring and initializing a variable.
	h := "apple"
	fmt.Println(h)

	// Variables can be reassigned, and several can be assigned at once, which makes
	// swapping easy.
	x, y := 1, 2
	x, y = y, x
	fmt.Println(x, y)

	// The blank identifier _ discards a value you do not need.
	_, second := 10, 20
	fmt.Println(second)

	// An unused local variable is a compile error, which keeps code tidy:
	// unused := 3
	// error: declared and not used: unused
}
