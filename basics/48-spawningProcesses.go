package main

import (
	"errors"
	"fmt"
	"io"
	"os/exec"
)

// Sometimes we need our program to spawn other processes
func spawningProcesses() {
	// Simple command
	dateCmd := exec.Command("date") // exec.Command creates an obj to represent the external process

	// Output runs the command, waits for it to complete, and returns the output
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
    fmt.Println(string(dateOut))

	// Output will return *exec.Error if the command fails (i.e. wrong path)
	// and *exec.ExitError if the command exits with a non-zero exit code
	_, err = exec.Command("date", "-x").Output()
	if err != nil {
		var execErr *exec.Error
		var exitErr *exec.ExitError
		switch {
		case errors.As(err, &execErr):
            fmt.Println("failed executing:", err)
		case errors.As(err, &exitErr):
            exitCode := exitErr.ExitCode()
            fmt.Println("command exit rc =", exitCode)
		default:
			panic(err)
		}
	}

	// More complex case
	// Pipe date to the external cmd's stdin, collect results from its stdout
	grepCmd := exec.Command("grep", "hello")

	grepIn, _ := grepCmd.StdinPipe()
    grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start() // Start the command
    grepIn.Write([]byte("hello grep\ngoodbye grep")) // Write to the stdin
	grepIn.Close()
	grepBytes, _ := io.ReadAll(grepOut) // Read from the stdout
	grepCmd.Wait() // Wait for the command to finish

	fmt.Println("> grep hello")
    fmt.Println(string(grepBytes))

	// When spawning a cmd, we need to provide 
	// an explicitly delineated command and argument array,
	// vs. being able to just pass in one command-line string.
	// We can use -c to spawn from a string
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
    lsOut, err := lsCmd.Output()
	if err != nil {
        panic(err)
    }
    fmt.Println("> ls -a -l -h")
    fmt.Println(string(lsOut))
}
