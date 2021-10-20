package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		time.Sleep(time.Second)
		ch <- 42
	}()
	fmt.Println(<-ch)
}
