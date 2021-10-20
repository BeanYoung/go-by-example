package main

import "fmt"

func fib(i int) int {
	switch i {
	case 1, 2:
		return 1
	default:
		return fib(i-1) + fib(i-2)
	}
}

func main() {
	fmt.Println(fib(11))
}
