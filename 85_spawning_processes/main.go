// Spawning Processes
//
// Sometimes our Go programs need to spawn other processes. The os/exec package runs
// external commands, wires up their input and output, and reports how they exited.
//
// Run: go run ./85_spawning_processes
package main

import (
	"errors"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

func main() {
	// We'll start with a simple command that takes no arguments or input and just
	// prints something to stdout. The exec.Command helper creates an object to
	// represent this external process.
	dateCmd := exec.Command("echo", "hello from echo")

	// The Output method runs the command, waits for it to finish and collects its
	// standard output. If there were no errors, dateOut will hold bytes with the
	// output.
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Print("> echo\n", string(dateOut))

	// Output and other methods of Command will return *exec.Error if there was a
	// problem executing the command (e.g. wrong path), and *exec.ExitError if the
	// command ran but exited with a non-zero return code.
	_, err = exec.Command("ls", "/a/nonexistent/path").Output()
	var execErr *exec.Error
	var exitErr *exec.ExitError
	switch {
	case errors.As(err, &execErr):
		fmt.Println("failed executing:", err)
	case errors.As(err, &exitErr):
		// The exact code differs between systems, so we only check it is non-zero.
		fmt.Println("command failed with a non-zero exit code:", exitErr.ExitCode() != 0)
	default:
		panic(err)
	}

	// LookPath reports whether a program can be found on the PATH.
	_, err = exec.LookPath("no-such-program-xyz")
	fmt.Println("lookpath:", errors.Is(err, exec.ErrNotFound))

	// Next we'll look at a slightly more involved case where we pipe data to the
	// external process on its stdin and collect the results from its stdout.
	grepCmd := exec.Command("grep", "hello")

	// Here we explicitly grab input/output pipes, start the process, write some
	// input to it, read the resulting output, and finally wait for the process to
	// exit.
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := io.ReadAll(grepOut)
	grepCmd.Wait()

	// We omitted error checks in the above example, but you could use the usual if
	// err != nil pattern for all of them. We also only collect the StdoutPipe
	// results, but you could collect the StderrPipe in exactly the same way.
	fmt.Println("> grep hello")
	fmt.Print(string(grepBytes))

	// Note that when spawning commands we need to provide an explicitly delineated
	// command and argument array, vs. being able to just pass in one command-line
	// string. If you want to spawn a full command with a string, you can use bash's
	// -c option.
	lsCmd := exec.Command("bash", "-c", "echo one two three | wc -w")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> wc -w")
	fmt.Println(strings.TrimSpace(string(lsOut)))

	// Stdin can also be any io.Reader, and CombinedOutput collects stdout and
	// stderr together.
	sortCmd := exec.Command("sort")
	sortCmd.Stdin = strings.NewReader("pear\napple\nfig\n")
	sorted, _ := sortCmd.CombinedOutput()
	fmt.Print("> sort\n", string(sorted))
}
