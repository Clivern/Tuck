// Readers and Writers
//
// io.Reader and io.Writer are the two most important interfaces in the standard
// library. Files, network connections, buffers, compressors, and hashes all
// implement them, so code written against them works with all of these.
//
// Run: go run ./69_readers_and_writers
package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"
)

// upperWriter is our own Writer: it upper-cases everything written to it before
// passing it on.
type upperWriter struct {
	w io.Writer
}

func (u upperWriter) Write(p []byte) (int, error) {
	return u.w.Write(bytes.ToUpper(p))
}

func main() {
	// A Reader has one method, Read(p []byte) (n int, err error). It fills p and
	// returns io.EOF when the data runs out.
	r := strings.NewReader("hello, reader")
	buf := make([]byte, 5)
	for {
		n, err := r.Read(buf)
		fmt.Printf("%d %q\n", n, buf[:n])
		if err == io.EOF {
			break
		}
	}

	// io.ReadAll reads everything; io.Copy streams from a Reader to a Writer.
	all, _ := io.ReadAll(strings.NewReader("read it all"))
	fmt.Println(string(all))
	io.Copy(os.Stdout, strings.NewReader("copied to stdout\n"))

	// bytes.Buffer is both a Reader and a Writer.
	var b bytes.Buffer
	b.WriteString("buffered ")
	fmt.Fprintf(&b, "%d bytes", 14)
	fmt.Println(b.String(), b.Len())

	// Our upperWriter works anywhere a Writer is expected.
	fmt.Fprintln(upperWriter{os.Stdout}, "shouting through a writer")

	// Readers and writers compose. MultiWriter writes to several writers, and
	// TeeReader copies what it reads into a writer.
	var copy1, copy2 bytes.Buffer
	io.WriteString(io.MultiWriter(&copy1, &copy2), "twice")
	fmt.Println(copy1.String(), copy2.String())

	var seen bytes.Buffer
	io.ReadAll(io.TeeReader(strings.NewReader("teed"), &seen))
	fmt.Println(seen.String())

	// LimitReader stops after n bytes, and MultiReader joins readers end to end.
	limited, _ := io.ReadAll(io.LimitReader(strings.NewReader("0123456789"), 4))
	joined, _ := io.ReadAll(io.MultiReader(strings.NewReader("ab"), strings.NewReader("cd")))
	fmt.Println(string(limited), string(joined))

	// A gzip.Writer compresses whatever passes through it, and a gzip.Reader
	// decompresses. Same interfaces, more power.
	var zipped bytes.Buffer
	zw := gzip.NewWriter(&zipped)
	io.WriteString(zw, strings.Repeat("go ", 100))
	zw.Close()
	fmt.Println("compressed 300 bytes to", zipped.Len())

	zr, _ := gzip.NewReader(&zipped)
	plain, _ := io.ReadAll(zr)
	fmt.Println("back to", len(plain), "bytes")

	// io.Pipe connects a writer in one goroutine to a reader in another.
	pr, pw := io.Pipe()
	go func() {
		fmt.Fprint(pw, "through a pipe")
		pw.Close()
	}()
	piped, _ := io.ReadAll(pr)
	fmt.Println(string(piped))
}
