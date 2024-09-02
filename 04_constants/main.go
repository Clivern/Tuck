// Constants
//
// Go supports constants of character, string, boolean, and numeric values. Constants
// are evaluated at compile time and cannot change.
//
// Run: go run ./04_constants
package main

import (
	"fmt"
	"math"
)

// const declares a constant value. It can appear anywhere a var can.
const s string = "constant"

// iota counts up from zero inside a const block, which makes enumerations easy.
const (
	small = iota
	medium
	large
)

func main() {
	fmt.Println(s)

	// A numeric constant has arbitrary precision until it is given a type.
	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)

	// An untyped constant takes the type its context needs, here an int64 and then
	// the float64 that math.Sin expects.
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))

	fmt.Println(small, medium, large)

	// Constant arithmetic is exact, unlike float variables.
	a, b := 0.1, 0.2
	fmt.Println(0.1+0.2 == 0.3, a+b == 0.3)
}
