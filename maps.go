package main

import "fmt"

func main() {
	m := map[string]int{}
	m["k1"] = 1
	m["k2"] = 2
	fmt.Println(m)

	delete(m, "k1")
	fmt.Println(m)
	_, presitent := m["k1"]
	fmt.Println(presitent)
}
