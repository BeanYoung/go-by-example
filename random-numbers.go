package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Float64())

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	fmt.Println(r.Intn(100))
}
