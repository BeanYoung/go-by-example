package main

import "fmt"

func main() {
	ch := make(chan string, 5)
	done := make(chan bool, 1)

	go func() {
		for {
			msg, more := <-ch
			if more {
				fmt.Println(msg)
			} else {
				done <- true
			}
		}
	}()

	for i := 0; i < 3; i++ {
		ch <- fmt.Sprintf("msg%d", i)
	}
	close(ch)
	<-done
}
