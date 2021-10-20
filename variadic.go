package main

import "fmt"

func sum(nums ...int) int {
	fmt.Println(nums)
	s := 0
	for _, n := range nums {
		s += n
	}
	return s
}

func main() {
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(3, 2, 2))
}
