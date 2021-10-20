package main

import "fmt"

func foo(i int) {
	i = 0
}

func bar(i *int) {
	*i = 0
}

func main() {
	i := 10
	foo(i)
	fmt.Println(i)
	bar(&i)
	fmt.Println(i)
}
