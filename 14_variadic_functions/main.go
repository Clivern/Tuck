// Variadic Functions
//
// A variadic function can be called with any number of trailing arguments.
// fmt.Println is a common one.
//
// Run: go run ./14_variadic_functions
package main

import (
	"fmt"
	"strings"
)

// This function takes any number of ints. Inside it, nums is a []int.
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

// Only the last parameter can be variadic, and it can follow regular ones.
func join(sep string, parts ...string) string {
	return strings.Join(parts, sep)
}

func main() {
	// Call it with individual arguments, or none at all.
	sum(1, 2)
	sum(1, 2, 3)
	sum()

	// If you already have the values in a slice, spread it with slice....
	nums := []int{1, 2, 3, 4}
	sum(nums...)

	fmt.Println(join("-", "a", "b", "c"))

	// any as the element type accepts values of every type, like fmt.Println does.
	show := func(vals ...any) {
		for _, v := range vals {
			fmt.Printf("%v(%T) ", v, v)
		}
		fmt.Println()
	}
	show(1, "two", 3.0)
}
