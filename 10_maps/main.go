// Maps
//
// Maps are Go's built-in associative data type, called hashes or dicts in other
// languages. The maps package adds helpers for them.
//
// Run: go run ./10_maps
package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	// Create an empty map with make(map[key-type]value-type).
	m := make(map[string]int)

	// Set key/value pairs with name[key] = val.
	m["k1"] = 7
	m["k2"] = 13

	// Printing a map shows all of its pairs, with keys sorted.
	fmt.Println("map:", m)

	// Get a value with name[key].
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// A missing key returns the zero value of the value type.
	v3 := m["k3"]
	fmt.Println("v3:", v3)

	// len returns the number of pairs.
	fmt.Println("len:", len(m))

	// delete removes a pair, and clear removes all of them.
	delete(m, "k2")
	fmt.Println("map:", m)
	clear(m)
	fmt.Println("map:", m)

	// The optional second result tells you whether the key was present. It
	// separates a missing key from a key holding a zero value.
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// Declare and initialize a map in one line.
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	// Iteration order over a map is not specified. Sort the keys when the order
	// matters.
	for _, k := range slices.Sorted(maps.Keys(n)) {
		fmt.Println(k, n[k])
	}

	// The maps package has helpers such as Equal and Clone.
	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}

	// A map of struct{} is the idiomatic set.
	seen := map[string]struct{}{}
	for _, w := range []string{"a", "b", "a"} {
		seen[w] = struct{}{}
	}
	fmt.Println("unique:", len(seen))

	// Writing to a nil map panics; reading from one is fine.
	var nilMap map[string]int
	fmt.Println("nil read:", nilMap["x"])
}
