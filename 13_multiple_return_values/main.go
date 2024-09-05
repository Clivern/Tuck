// Multiple Return Values
//
// Go has built-in support for multiple return values. This is used throughout
// idiomatic Go, for example to return both a result and an error.
//
// Run: go run ./13_multiple_return_values
package main

import (
	"errors"
	"fmt"
)

// (int, int) in the signature says the function returns two ints.
func vals() (int, int) {
	return 3, 7
}

// The last result is conventionally an error, nil on success.
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Results can be named. They start at their zero values and a bare return sends them
// back. This suits short functions, where the names document the results.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	// Use the different results with multiple assignment.
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// If you only want some of the values, use the blank identifier _.
	_, c := vals()
	fmt.Println(c)

	// Check the error before using the result.
	if q, err := divide(1, 0); err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println(q)
	}

	fmt.Println(split(17))
}
