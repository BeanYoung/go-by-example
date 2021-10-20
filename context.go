package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println("start process request")
	ctx := req.Context()
	select {
	case <-time.After(time.Second * 10):
		fmt.Fprintln(w, "hello")
		fmt.Println("10 seconds passed")
	case <-ctx.Done():
		fmt.Println("ctx done")
		fmt.Println(ctx.Err())
		err := ctx.Err()
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":9090", nil)
}
