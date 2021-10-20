package main

import (
	"fmt"
	"io"
	"os/exec"
)

func main() {
	dateCmd := exec.Command("date")
	fmt.Println(dateCmd.Output())

	grepCmd := exec.Command("grep", "hello")
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello, world\ndemo"))
	grepIn.Close()
	fmt.Println(io.ReadAll(grepOut))
	grepCmd.Wait()
}
