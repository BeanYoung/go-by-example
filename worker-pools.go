package main

import (
	"fmt"
	"time"
)

func worker(i int, jobs <-chan int, results chan<- int) {
	fmt.Println("start worker", i)
	for j := range jobs {
		fmt.Println("start job", j)
		time.Sleep(time.Second * 2)
		fmt.Println("finish job", j)
		results <- j * 2
	}
	fmt.Println("finish worker", i)
}

func main() {
	workerNum := 5
	jobs := make(chan int, workerNum)
	results := make(chan int, workerNum)

	for i := 0; i < workerNum; i++ {
		go worker(i, jobs, results)
	}

	for i := 0; i < 10; i++ {
		jobs <- i
	}

	close(jobs)

	for j := 0; j < 10; j++ {
		fmt.Println(<-results)
	}
}
