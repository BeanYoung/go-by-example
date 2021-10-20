package main

import "fmt"

func main() {
	i := 0
	for {
		fmt.Println(i)
		if i > 10 {
			break
		}
		i++
	}

	for j := 0; j < 10; j++ {
		fmt.Println(j)
	}

	k := 0
	for k < 10 {
		fmt.Println(k)
		k++
	}
}
