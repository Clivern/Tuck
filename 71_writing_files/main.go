// Writing Files
//
// Writing files in Go follows similar patterns to the ones we saw earlier for
// reading.
//
// Run: go run ./71_writing_files
package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dir, err := os.MkdirTemp("", "writing")
	check(err)
	defer os.RemoveAll(dir)

	// To start, here's how to dump a string (or just bytes) into a file. 0644 is the
	// file mode: read and write for the owner, read for everyone else.
	d1 := []byte("hello\ngo\n")
	err = os.WriteFile(filepath.Join(dir, "dat1"), d1, 0644)
	check(err)

	// For more granular writes, open a file for writing.
	f, err := os.Create(filepath.Join(dir, "dat2"))
	check(err)

	// It's idiomatic to defer a Close immediately after opening a file.
	defer f.Close()

	// You can Write byte slices as you'd expect.
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// A WriteString is also available.
	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	// Issue a Sync to flush writes to stable storage.
	f.Sync()

	// bufio provides buffered writers in addition to the buffered readers we saw
	// earlier.
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	// Use Flush to ensure all buffered operations have been applied to the
	// underlying writer.
	w.Flush()

	// OpenFile with flags appends to an existing file instead of truncating it.
	af, err := os.OpenFile(filepath.Join(dir, "dat1"), os.O_APPEND|os.O_WRONLY, 0644)
	check(err)
	fmt.Fprintln(af, "appended")
	af.Close()

	// Read both files back to see what we wrote.
	for _, name := range []string{"dat1", "dat2"} {
		data, err := os.ReadFile(filepath.Join(dir, name))
		check(err)
		fmt.Printf("%s:\n%s", name, data)
	}
}
