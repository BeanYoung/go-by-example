package main

import (
	"bufio"
	"os"
)

func main() {
	os.WriteFile("/tmp/data1", []byte("demo"), 0664)

	f, _ := os.Create("/tmp/data2")
	defer f.Close()
	f.WriteString("demo")

	f1, _ := os.Create("/tmp/data2")
	defer f1.Close()
	writer := bufio.NewWriter(f1)
	writer.WriteString("demo")
	writer.Flush()
}
