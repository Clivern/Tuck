// Sorting by Functions
//
// Sometimes we'll want to sort a collection by something other than its natural
// order. For example, suppose we wanted to sort strings by their length instead of
// alphabetically. Here's an example of custom sorts in Go.
//
// Run: go run ./53_sorting_by_functions
package main

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
)

func main() {
	fruits := []string{"peach", "banana", "kiwi"}

	// We implement a comparison function for string lengths. cmp.Compare is helpful
	// for this: it returns a negative number, zero, or a positive number.
	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	// Now we can call slices.SortFunc with this custom comparison function to sort
	// fruits by name length.
	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)

	// We can use the same technique to sort a slice of values that aren't built-in
	// types.
	type Person struct {
		name string
		age  int
	}

	people := []Person{
		Person{name: "Jax", age: 37},
		Person{name: "TJ", age: 25},
		Person{name: "Alex", age: 72},
		Person{name: "Bea", age: 25},
	}

	// Sort people by age using slices.SortFunc.
	slices.SortFunc(people,
		func(a, b Person) int {
			return cmp.Compare(a.age, b.age)
		})
	fmt.Println(people)

	// cmp.Or chains comparisons: by age, then by name for equal ages. Swapping a and
	// b sorts in descending order.
	slices.SortFunc(people, func(a, b Person) int {
		return cmp.Or(
			cmp.Compare(b.age, a.age),
			strings.Compare(a.name, b.name),
		)
	})
	fmt.Println(people)

	// SortStableFunc keeps equal elements in their original order.
	words := []string{"bb", "a", "cc", "d"}
	slices.SortStableFunc(words, lenCmp)
	fmt.Println(words)

	// Note: if the Person struct is large, you may want the slice to contain
	// *Person instead and adjust the sorting function accordingly. If in doubt,
	// benchmark!
}
