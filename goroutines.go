package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second)
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("sync")
	go f("async")
	go func(m string) {
		f(m)
	}("async lambda")
	time.Sleep(time.Second * 4)
}
