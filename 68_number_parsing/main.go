// Number Parsing
//
// Parsing numbers from strings is a basic but common task in many programs; here's
// how to do it in Go.
//
// Run: go run ./68_number_parsing
package main

// The built-in package strconv provides the number parsing.
import (
	"fmt"
	"strconv"
)

func main() {
	// With ParseFloat, this 64 tells how many bits of precision to parse.
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// For ParseInt, the 0 means infer the base from the string. 64 requires that
	// the result fit in 64 bits.
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	// ParseInt will recognize hex-formatted numbers.
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// A ParseUint is also available.
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// Atoi is a convenience function for basic base-10 int parsing.
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	// Parse functions return an error on bad input.
	_, e := strconv.Atoi("wat")
	fmt.Println(e)

	// Values that do not fit are reported too.
	_, e = strconv.ParseInt("300", 10, 8)
	fmt.Println(e)

	// ParseBool accepts 1, t, T, TRUE, true, True, and their false counterparts.
	b, _ := strconv.ParseBool("true")
	fmt.Println(b)

	// Going the other way: Itoa, FormatInt with a base, FormatFloat, and Quote.
	fmt.Println(strconv.Itoa(42), strconv.FormatInt(255, 2), strconv.FormatFloat(3.14159, 'f', 2, 64))
	fmt.Println(strconv.Quote("tab\there"))
}
