// Recursion
//
// Go supports recursive functions. Here's a classic example.
//
// Run: go run ./16_recursion
package main

import "fmt"

// fact calls itself until it reaches the base case, fact(0).
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))

	// An anonymous function can be recursive too, but the variable must be declared
	// with var before the function is defined, so the body can refer to it.
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))

	// Recursion fits recursive data such as trees.
	type node struct {
		val         int
		left, right *node
	}
	var sum func(n *node) int
	sum = func(n *node) int {
		if n == nil {
			return 0
		}
		return n.val + sum(n.left) + sum(n.right)
	}
	tree := &node{1, &node{2, nil, nil}, &node{3, &node{4, nil, nil}, nil}}
	fmt.Println(sum(tree))
}
