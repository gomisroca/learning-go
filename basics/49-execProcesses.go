package main

import (
	"os"
	"os/exec"
	"syscall"
)

// Sometimes we want to completely replace the current process with another
// Go manages this with exec
func execProcesses() {
	// We will exec ls
	// Go requires an absolute path to the executable
	binary, lookErr := exec.LookPath("ls") // LookPath searches the PATH for the executable
	if lookErr != nil {
		panic(lookErr)
	}

	// exec requires args in slice form
	args := []string{"ls", "-a", "-l", "-h"}

	// it also requires an environment
	env := os.Environ()

	// now we do the syscall.Exec
	// if it is successful, the current process is replaced with the new one
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}