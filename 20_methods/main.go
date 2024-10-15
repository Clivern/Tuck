// Methods
//
// Go supports methods defined on struct types, and on any other named type declared
// in the same package.
//
// Run: go run ./20_methods
package main

import "fmt"

type rect struct {
	width, height int
}

// This area method has a receiver type of *rect.
func (r *rect) area() int {
	return r.width * r.height
}

// Methods can have pointer or value receivers. This one has a value receiver.
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

// A pointer receiver can change the value it is called on.
func (r *rect) scale(k int) {
	r.width *= k
	r.height *= k
}

// Methods can be declared on any named type, not only structs.
type celsius float64

func (c celsius) fahrenheit() float64 {
	return float64(c)*9/5 + 32
}

// An alias is another name for a type, not a new type, so it cannot declare
// methods. celsius can, because it is a defined type.
type fahrenheit = float64

// func (f fahrenheit) boiling() bool { return f >= 212 }
// error: cannot define new methods on non-local type fahrenheit

func main() {
	r := rect{width: 10, height: 5}

	// Call the two methods defined for our struct.
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	// Go converts between values and pointers for method calls automatically. Use a
	// pointer receiver to avoid copying on calls, or to let the method change the
	// struct.
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())

	r.scale(2)
	fmt.Println("scaled:", r)

	fmt.Println(celsius(100).fahrenheit())

	// The alias is float64, so values move between the two names with no conversion.
	var deg fahrenheit = celsius(100).fahrenheit()
	var plain float64 = deg
	fmt.Println("alias:", plain)

	// A method value binds the receiver, so it can be passed around as a function.
	perim := r.perim
	fmt.Println("method value:", perim())
}
