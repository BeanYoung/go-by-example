package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(strconv.ParseInt("123", 0, 64))
	fmt.Println(strconv.ParseFloat("0.111", 64))
	fmt.Println(strconv.Atoi("212"))
}
