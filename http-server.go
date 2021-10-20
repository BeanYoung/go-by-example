package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "hello")
	for k, v := range req.Header {
		fmt.Fprintf(w, "%s: %s", k, v)
	}
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":9090", nil)
}
