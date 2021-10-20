package main

import "fmt"

func Map(s []string) []int {
	resp := make([]int, len(s))
	for idx, v := range s {
		resp[idx] = len(v)
	}
	return resp
}

func main() {
	s := []string{"a", "ab", "abc"}
	fmt.Println(Map(s))
}
