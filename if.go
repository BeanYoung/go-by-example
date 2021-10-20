package main

import "fmt"

func main() {
	if true {
		fmt.Println(true)
	}

	if i := 10; i < 10 {
		fmt.Println("if")
	} else if i > 10 {
		fmt.Println("else if")
	} else {
		fmt.Println("else")
	}
}
