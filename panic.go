package main

import "os"

func main() {
	// panic("wops")

	_, err := os.Create("/tmp/tmp/log")
	if err != nil {
		panic(err)
	}
}
