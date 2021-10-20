package main

import "fmt"

func main() {
	s := [4]int{}
	fmt.Println(len(s))
	for idx, i := range s {
		fmt.Println(idx, i)
	}

	m := map[string]int{"k1": 1, "k2": 2}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
