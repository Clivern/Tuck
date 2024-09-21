// Random Numbers
//
// Go's math/rand/v2 package provides pseudorandom number generation. For anything
// security related, such as tokens and keys, use crypto/rand instead.
//
// Run: go run ./67_random_numbers
package main

import (
	crand "crypto/rand"
	"fmt"
	"math/rand/v2"
)

func main() {
	// For example, rand.IntN returns a random int n, 0 <= n < 100. The output
	// differs between runs, so we only check the range here.
	n := rand.IntN(100)
	fmt.Println("0 <= n < 100:", n >= 0 && n < 100)

	// rand.Float64 returns a float64 f, 0.0 <= f < 1.0.
	f := rand.Float64()
	fmt.Println("0 <= f < 1:", f >= 0 && f < 1)

	// This can be used to generate random floats in other ranges, for example
	// 5.0 <= f' < 10.0.
	g := rand.Float64()*5 + 5
	fmt.Println("5 <= g < 10:", g >= 5 && g < 10)

	// If you want a known seed, create a new rand.Source and pass it into the New
	// constructor. NewPCG creates a new PCG source that requires a seed of two
	// uint64 numbers. A seeded generator gives the same sequence every time.
	s2 := rand.NewPCG(42, 1024)
	r2 := rand.New(s2)
	fmt.Print(r2.IntN(100), ",")
	fmt.Print(r2.IntN(100))
	fmt.Println()

	s3 := rand.NewPCG(42, 1024)
	r3 := rand.New(s3)
	fmt.Print(r3.IntN(100), ",")
	fmt.Print(r3.IntN(100))
	fmt.Println()

	// Shuffle and Perm rearrange values; with a seeded generator they repeat too.
	words := []string{"a", "b", "c", "d", "e"}
	r4 := rand.New(rand.NewPCG(1, 2))
	r4.Shuffle(len(words), func(i, j int) { words[i], words[j] = words[j], words[i] })
	fmt.Println(words, r4.Perm(5))

	// crypto/rand reads from the operating system's secure source. rand.Text
	// returns a random string suitable for tokens.
	token := crand.Text()
	key := make([]byte, 16)
	crand.Read(key)
	fmt.Println("token length:", len(token), "key bytes:", len(key))
}
