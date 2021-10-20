package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 2)

	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			ch <- fmt.Sprintf("msg%d", i)
		}
		close(ch)
	}()

	for msg := range ch {
		fmt.Println(msg)
	}
}
