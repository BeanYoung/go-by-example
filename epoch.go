package main

import (
	"fmt"
	"time"
)

func main() {
	n := time.Now()
	timestamp := n.Unix()
	fmt.Println(timestamp)

	fmt.Println(n.UnixMicro())
	fmt.Println(n.UnixMilli())
	fmt.Println(n.UnixNano())
}
