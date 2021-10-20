package main

import (
	"errors"
	"fmt"
)

func foo(i int) (int, error) {
	if i == 42 {
		return -1, errors.New("i is 42")
	}
	return i + 3, nil
}

type httpError struct {
	status string
	code   int
}

func (self *httpError) Error() string {
	return fmt.Sprintf("http errorstatus: %s code: %d", self.status, self.code)
}

func parseHttp() error {
	return &httpError{"not_found", 404}
}

func main() {
	if i, err := foo(10); err != nil {
		fmt.Println("error", i)
	} else {
		fmt.Println("ok", i)
	}
	if i, err := foo(42); err != nil {
		fmt.Println("error", i)
	} else {
		fmt.Println("ok", i)
	}

	err := parseHttp()
	if he, ok := err.(*httpError); ok {
		fmt.Println(he.status)
		fmt.Println(he.code)
	}
}
