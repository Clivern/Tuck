// If/Else
//
// Branching with if and else is straightforward. Parentheses around the condition
// are not needed, but braces are.
//
// Run: go run ./06_if_else
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// A basic example.
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// An if without an else.
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// Logical operators combine conditions.
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 is even")
	}

	// A statement can come before the condition. Variables declared there are
	// available in the current and all later branches.
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// This is the most common shape in Go: call something, then check its error.
	if n, err := strconv.Atoi("42x"); err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("parsed", n)
	}

	// There is no ternary operator (a ? b : c) in Go; use a full if instead.
	// Conditions must be booleans, there is no truthiness:
	// if 1 {}
	// error: non-boolean condition in if statement
}
