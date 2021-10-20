package main

import (
	"flag"
	"fmt"
)

func main() {
	wordFlag := flag.String("word", "foo", "a string")
	numFlag := flag.Int("num", 0, "an int")
	forceFlag := flag.Bool("force", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "svar", "a string")

	flag.Parse()
	fmt.Println(*wordFlag, *numFlag, *forceFlag, svar, flag.Args())
}
