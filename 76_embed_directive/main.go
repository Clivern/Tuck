// Embed Directive
//
// //go:embed is a compiler directive that allows programs to include arbitrary files
// and folders in the Go binary at build time. Read more about the embed directive in
// the embed package documentation.
//
// Run: go run ./76_embed_directive
package main

// Import the embed package; if you don't use any exported identifiers from this
// package, you can do a blank import with _ "embed".
import (
	"embed"
	"fmt"
	"io/fs"
)

// embed directives accept paths relative to the directory containing the Go source
// file. This directive embeds the contents of the file into the string variable
// immediately following it.
//
//go:embed folder/single_file.txt
var fileString string

// Or embed the contents of the file into a []byte.
//
//go:embed folder/single_file.txt
var fileByte []byte

// We can also embed multiple files or even folders with wildcards. This uses a
// variable of the embed.FS type, which implements a simple virtual file system.
//
//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

func main() {
	// Print out the contents of single_file.txt.
	fmt.Print(fileString)
	fmt.Print(string(fileByte))

	// Retrieve some files from the embedded folder.
	content1, _ := folder.ReadFile("folder/file1.hash")
	fmt.Print(string(content1))

	content2, _ := folder.ReadFile("folder/file2.hash")
	fmt.Print(string(content2))

	// embed.FS implements fs.FS, so the io/fs helpers work on it.
	fs.WalkDir(folder, ".", func(path string, d fs.DirEntry, err error) error {
		fmt.Println("embedded:", path)
		return err
	})
}
