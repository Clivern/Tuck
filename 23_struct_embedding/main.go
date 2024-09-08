// Struct Embedding
//
// Go supports embedding of structs and interfaces to express a more seamless
// composition of types. Go has no inheritance; embedding is how types reuse each
// other.
//
// Run: go run ./23_struct_embedding
package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// A container embeds a base. An embedding looks like a field without a name.
type container struct {
	base
	str string
}

func main() {
	// When creating structs with literals, we initialize the embedding explicitly;
	// the embedded type serves as the field name.
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	// We can access the base's fields directly on co.
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// Alternatively, spell out the full path using the embedded type name.
	fmt.Println("also num:", co.base.num)

	// Since container embeds base, the methods of base also become methods of a
	// container. Here we call describe directly on co.
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	// Embedding structs with methods can bestow interface implementations onto
	// other structs. Here a container now implements the describer interface
	// because it embeds base.
	var d describer = co
	fmt.Println("describer:", d.describe())

	// Interfaces can embed other interfaces too, combining their methods.
	type namer interface{ name() string }
	type namedDescriber interface {
		describer
		namer
	}
	var _ namedDescriber = nil
	fmt.Println("interfaces compose")
}
