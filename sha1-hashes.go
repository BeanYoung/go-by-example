package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	s := sha1.New()
	bs := s.Sum([]byte("demo"))
	fmt.Printf("%x\n", bs)
}
