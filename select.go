package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		time.Sleep(time.Second)
		ch1 <- "msg1"
	}()
	go func() {
		time.Sleep(time.Second)
		ch2 <- "msg2"
	}()
	for i := 0; i < 2; i++ {
		select {
		case m := <-ch1:
			fmt.Println("ch1", m)
		case m := <-ch2:
			fmt.Println("ch2", m)
		}
	}
}
