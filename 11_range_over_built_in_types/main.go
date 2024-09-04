// Range over Built-in Types
//
// range iterates over elements in a variety of built-in data structures. Here is
// range with slices, maps, strings, and integers.
//
// Run: go run ./11_range_over_built_in_types
package main

import "fmt"

func main() {
	// Sum the numbers in a slice. Arrays work the same way.
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// range yields both the index and the value. Above we ignored the index with _.
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// The value is a copy; assign through the index to change the slice.
	for i := range nums {
		nums[i] *= 10
	}
	fmt.Println(nums)

	// range on a map iterates over key/value pairs.
	kvs := map[string]string{"a": "apple"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// range can also iterate over just the keys of a map.
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// range on a string iterates over Unicode code points. The first value is the
	// byte index of the rune, the second the rune itself.
	for i, c := range "go€" {
		fmt.Println(i, c, string(c))
	}

	// range over an integer counts from zero.
	for i := range 3 {
		fmt.Print(i, " ")
	}
	fmt.Println()
}
