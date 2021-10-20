package main

import (
	"fmt"
	"os"
)

func main() {
	os.Setenv("BAR", "1")
	fmt.Println(os.Getenv("FOO"))
	fmt.Println(os.Getenv("BAR"))

	for k, v := range os.Environ() {
		fmt.Println(k, v)
	}
}
