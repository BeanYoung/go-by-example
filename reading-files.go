package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	d, _ := os.ReadFile("/tmp/data")
	fmt.Println(string(d))

	f, _ := os.Open("/tmp/data")
	b := make([]byte, 2)
	n, _ := f.Read(b)
	fmt.Println(n, string(b))

	f.Seek(1, 0)
	fmt.Println(f.Read(b))

	io.ReadAtLeast(f, b, 2)
	fmt.Println(string(b))

	reader := bufio.NewReader(f)
	fmt.Println(reader.Peek(5))
}
