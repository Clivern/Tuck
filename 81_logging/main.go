// Logging
//
// The Go standard library provides straightforward tools for outputting logs from Go
// programs, with the log package for free-form output and the log/slog package for
// structured output.
//
// Run: go run ./81_logging
package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	"log/slog"
)

func main() {
	// Simply invoking functions like Println from the log package uses the standard
	// logger, which is already pre-configured for reasonable logging output to
	// os.Stderr. Additional methods like Fatal* or Panic* will exit the program
	// after logging. We turn off the timestamp here so the output is stable.
	log.SetFlags(0)
	log.SetOutput(os.Stdout)
	log.Println("standard logger")

	// Loggers can be configured with flags to set their output format. The standard
	// logger has the log.LstdFlags flag set by default, which prints the date and
	// time. Lshortfile adds the file name and line number.
	log.SetFlags(log.Lshortfile)
	log.Println("with file/line")

	// It may be useful to create a custom logger and pass it around. When creating
	// a new logger, we can set a prefix to distinguish its output from other
	// loggers.
	mylog := log.New(os.Stdout, "my:", 0)
	mylog.Println("from mylog")

	// We can set the prefix on existing loggers (including the standard one) with
	// the SetPrefix method.
	mylog.SetPrefix("ohmy:")
	mylog.Println("from mylog")

	// Loggers can have custom output targets; any io.Writer works.
	var buf bytes.Buffer
	buflog := log.New(&buf, "buf:", 0)

	// This call writes the log output into buf.
	buflog.Println("hello")

	// This will actually show it on standard output.
	fmt.Print("from buflog:", buf.String())

	// The slog package provides structured log output. For example, logging in JSON
	// format is straightforward. We drop the time attribute to keep the output
	// stable.
	noTime := func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.TimeKey {
			return slog.Attr{}
		}
		return a
	}
	jsonHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{ReplaceAttr: noTime})
	myslog := slog.New(jsonHandler)
	myslog.Info("hi there")

	// In addition to the message, slog output can contain an arbitrary number of
	// key=value pairs.
	myslog.Info("hello again", "key", "val", "age", 25)

	// The text handler prints key=value pairs. With adds attributes to every record,
	// and Level filters out less important ones.
	text := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:       slog.LevelWarn,
		ReplaceAttr: noTime,
	})).With("service", "api")
	text.Info("not shown, below the level")
	text.Warn("disk almost full", "free", "2%", "path", strings.Repeat("/x", 2))
	text.Error("request failed", slog.Group("req", "method", "GET", "status", 500))
}
