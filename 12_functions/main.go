// Functions
//
// Functions are central in Go. We'll learn about them through a few examples.
//
// Run: go run ./12_functions
package main

import "fmt"

// This function takes two ints and returns their sum as an int. Parameter types come
// after the names, and the result type after the parameter list.
func plus(a int, b int) int {
	// Go requires explicit returns: it does not return the value of the last
	// expression automatically.
	return a + b
}

// Consecutive parameters of the same type can share one type name.
func plusPlus(a, b, c int) int {
	return a + b + c
}

// Arguments are passed by value: the function gets a copy, so this does not change
// the caller's variable. Pointers (see Pointers) change that.
func reset(n int) {
	n = 0
}

func main() {
	// Call a function with name(args).
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	n := 5
	reset(n)
	fmt.Println("still", n)

	// Functions are values too: they can be assigned and passed around.
	f := plus
	fmt.Printf("%T %d\n", f, f(10, 20))
}
