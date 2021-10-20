package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println(a)
	a[4] = 10
	fmt.Println(a)
	fmt.Println("len(a)", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	var c []int
	c = append(c, 1)
	fmt.Println(c)

	var d [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			d[i][j] = i * j
		}
	}
	fmt.Println(d)
}
