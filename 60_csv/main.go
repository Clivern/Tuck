// CSV
//
// The encoding/csv package reads and writes comma-separated values, as described in
// RFC 4180. It handles quoting, embedded commas, and newlines inside fields.
//
// Run: go run ./60_csv
package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const data = `name,lang,stars
tuck,Go,120
"buzzard, the rust one",Rust,95
nut,"CSS",40
`

func main() {
	// A csv.Reader reads from any io.Reader. ReadAll returns every record as a
	// []string.
	r := csv.NewReader(strings.NewReader(data))
	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	for _, rec := range records {
		fmt.Printf("%q\n", rec)
	}

	// For large inputs, read one record at a time until io.EOF.
	r = csv.NewReader(strings.NewReader(data))
	header, _ := r.Read()
	fmt.Println("columns:", header)
	total := 0
	for {
		rec, err := r.Read()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			panic(err)
		}
		n, _ := strconv.Atoi(rec[2])
		total += n
	}
	fmt.Println("total stars:", total)

	// The reader can be tuned: a different separator, comments, and a check that
	// every record has the same number of fields.
	tsv := csv.NewReader(strings.NewReader("# comment\na\tb\nc\td\te\n"))
	tsv.Comma = '\t'
	tsv.Comment = '#'
	_, err = tsv.ReadAll()
	fmt.Println("error:", err)

	// A csv.Writer quotes fields as needed. Call Flush when done and check Error.
	w := csv.NewWriter(os.Stdout)
	w.Write([]string{"id", "note"})
	w.Write([]string{"1", "plain"})
	w.Write([]string{"2", `has "quotes", and a comma`})
	w.Flush()
	if err := w.Error(); err != nil {
		panic(err)
	}
}
