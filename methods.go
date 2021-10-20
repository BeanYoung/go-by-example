package main

import "fmt"

type rect struct {
	x int
	y int
}

func (self *rect) area() int {
	return self.x * self.y
}

func main() {
	r := rect{10, 20}
	fmt.Println(r.area())

	rp := &rect{10, 20}
	fmt.Println(rp.area())
}
