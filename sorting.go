package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"c", "b", "a"}
	sort.Strings(s)
	fmt.Println(s)

	i := []int{3, 2, 1}
	sort.Ints(i)
	fmt.Println(i)

	fmt.Println(sort.IntsAreSorted(i))
}
