package main

import "fmt"

func nextInt() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextIntFunc := nextInt()
	fmt.Println(nextIntFunc())
	fmt.Println(nextIntFunc())
	fmt.Println(nextIntFunc())
	fmt.Println(nextIntFunc())
}
