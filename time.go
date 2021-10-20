package main

import (
	"fmt"
	"time"
)

func main() {
	n := time.Now()
	fmt.Println(n)

	t := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println(t)

	fmt.Println(n.Year())

	fmt.Println(n.Sub(t))

	fmt.Println(n.Add(time.Hour * 2))
}
