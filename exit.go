package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		fmt.Println("done")
	}()
	os.Exit(3)
}
