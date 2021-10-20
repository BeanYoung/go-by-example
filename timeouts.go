package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		ch1 <- "msg1"
	}()
	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case <-time.After(time.Second):
		fmt.Println("ch1 timeout")
	}
}
