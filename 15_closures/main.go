// Closures
//
// Go supports anonymous functions, which can form closures. Anonymous functions are
// useful when you want to define a function inline without having to name it.
//
// Run: go run ./15_closures
package main

import (
	"fmt"
	"slices"
	"strings"
)

// intSeq returns another function, defined anonymously in its body. The returned
// function closes over the variable i to form a closure.
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// nextInt captures its own i, which is updated each time we call it.
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// The state is unique to each closure: a new one starts over.
	newInts := intSeq()
	fmt.Println(newInts())

	// Closures capture variables, not values. A change after the closure is made is
	// visible inside it.
	x := 1
	show := func() { fmt.Println("x is", x) }
	x = 2
	show()

	// Closures are handy as arguments, for example to sort by a custom rule.
	words := []string{"banana", "Apple", "cherry"}
	slices.SortFunc(words, func(a, b string) int {
		return strings.Compare(strings.ToLower(a), strings.ToLower(b))
	})
	fmt.Println(words)

	// A function literal can be called right where it is defined.
	fmt.Println(func(a, b int) int { return a * b }(6, 7))
}
