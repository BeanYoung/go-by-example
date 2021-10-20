package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		requests <- i
	}
	close(requests)
	tick := time.Tick(time.Millisecond * 200)
	for request := range requests {
		<-tick
		fmt.Println("Request ", request)
	}

	busyLimiter := make(chan time.Time, 2)
	for i := 0; i < 2; i++ {
		busyLimiter <- time.Now()
	}
	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			busyLimiter <- t
		}
	}()

	requests1 := make(chan int, 5)
	for i := 0; i < 5; i++ {
		requests1 <- i
	}
	close(requests1)
	for request := range requests1 {
		<-busyLimiter
		fmt.Println("Busy limit request", request)
	}
}
