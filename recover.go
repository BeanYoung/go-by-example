package main

import "fmt"

func manyPanic() {
	panic("doom")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	manyPanic()
	fmt.Println("never goes here")
}
