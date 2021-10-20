package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(time.Second * 2)

	go func() {
		<-timer1.C
		fmt.Println("time1 fired")
	}()

	if time1Stopped := timer1.Stop(); time1Stopped {
		fmt.Println("timer1 stopped")
	}
	time.Sleep(time.Second * 3)
}
