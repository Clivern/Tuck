// Values
//
// Go has strings, integers, floats, booleans, and more. Here are a few basic values
// and the operators that work on them.
//
// Run: go run ./02_values
package main

import "fmt"

func main() {
	// Strings can be joined with +.
	fmt.Println("go" + "lang")

	// Integers and floats.
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// Integer division truncates; % is the remainder.
	fmt.Println("7/3 =", 7/3, "7%3 =", 7%3)

	// Booleans, with the usual boolean operators.
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	// Integers have explicit widths. A typed number does not convert to another
	// width on its own; a conversion names the type you want.
	var tiny int8 = 100
	fmt.Println("widths:", tiny, int16(tiny), byte('A'))

	// Bitwise operators act on integers: and, or, xor, and-not, and shifts.
	fmt.Println("bits:", 0b1100&0b1010, 0b1100|0b1010, 0b1100^0b1010, 0b1100&^0b0100, 1<<3, 8>>2)

	// A complex number holds two floats. real and imag take them back out.
	z := 3 + 4i
	fmt.Println("complex:", real(z), imag(z), complex(1, 2))
}
