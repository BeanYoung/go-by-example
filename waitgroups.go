package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(i int) {
	fmt.Printf("Worker %d start\n", i)
	time.Sleep(time.Second * time.Duration(i))
	fmt.Printf("Worker %d done\n", i)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		i := i

		go func() {
			defer wg.Done()
			worker(i)
		}()
	}
	wg.Wait()
}
