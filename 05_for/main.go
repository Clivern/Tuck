// For
//
// for is Go's only looping construct. Here are its basic forms.
//
// Run: go run ./05_for
package main

import "fmt"

func main() {
	// The most basic form has a single condition, like a while loop.
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// The classic initial; condition; after loop.
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// range over an integer counts from 0 up to, but not including, the number.
	for i := range 3 {
		fmt.Println("range", i)
	}

	// for without a condition loops until you break out of it or return.
	for {
		fmt.Println("loop")
		break
	}

	// continue skips to the next iteration.
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	// A label lets break and continue act on an outer loop.
outer:
	for i := range 3 {
		for j := range 3 {
			if i+j == 3 {
				break outer
			}
			fmt.Println(i, j)
		}
	}

	// goto jumps to a label in the same function. A labeled break is usually
	// clearer. goto cannot skip the declaration of a variable.
	n := 0
jump:
	if n == 2 {
		goto finished
	}
	fmt.Println("goto", n)
	n++
	goto jump
finished:
	fmt.Println("finished")
}
