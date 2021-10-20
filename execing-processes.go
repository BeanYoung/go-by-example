package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	binary, err := exec.LookPath("ls")
	if err != nil {
		panic("ls not found")
	}
	args := []string{"ls", "-a", "-s"}
	env := os.Environ()
	err = syscall.Exec(binary, args, env)
	fmt.Println(err)
}
