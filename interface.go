package main

import "fmt"

type animal interface {
	run(distance int) int
	eat() string
}

type cat struct {
	speed int
	size  int
}

type dog struct {
	speed  int
	weight int
}

func (self cat) run(distance int) int {
	return distance / self.speed
}

func (self cat) eat() string {
	return "cat eat fish"
}

func (self dog) run(distance int) int {
	return distance / self.speed * 4
}

func (self dog) eat() string {
	return "dog eat shit"
}

func inspect(a animal) {
	fmt.Println(a.eat())
	fmt.Println(a.run(100))
}

func main() {
	c := cat{10, 20}
	d := dog{30, 40}
	inspect(c)
	inspect(d)
}
