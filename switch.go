package main

import (
	"fmt"
	"time"
)

func whoami(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("int")
	case bool:
		fmt.Println("bool")
	default:
		fmt.Println("others")
	}
}

func main() {
	i := 10
	switch i {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("default")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("weekend")
	default:
		fmt.Println("weekday")
	}
	whoami(1)
	whoami(true)
	whoami("")
}
