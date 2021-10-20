package main

import "fmt"

type Person struct {
	name string
	age  int
}

func newPerson(name string) *Person {
	return &Person{name, 42}
}

func main() {
	person := Person{"demo", 1}
	fmt.Println(person)
	pointer := &Person{"pointer", 10}
	fmt.Println(pointer)

	p1 := newPerson("new person")
	p1.age = 50
	fmt.Println(p1)
}
