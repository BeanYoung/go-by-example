package main

import "fmt"

func main() {
	var s []string
	s = append(s, "1")
	fmt.Println(s, len(s))

	t := make([]string, 5)
	t[2] = "2"
	fmt.Println(t, len(t))
}
