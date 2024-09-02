// Values
//
// Go has strings, integers, floats, booleans, and more. Here are a few basic values
// and the operators that work on them.
//
// Run: go run ./02_values
package main

import "fmt"

func main() {
	// Strings can be joined with +.
	fmt.Println("go" + "lang")

	// Integers and floats.
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// Integer division truncates; % is the remainder.
	fmt.Println("7/3 =", 7/3, "7%3 =", 7%3)

	// Booleans, with the usual boolean operators.
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
