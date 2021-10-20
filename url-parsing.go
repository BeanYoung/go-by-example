package main

import (
	"fmt"
	"net/url"
)

func main() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"
	u, _ := url.Parse(s)
	fmt.Println(u)
	fmt.Println(u.Scheme)
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)
	fmt.Println(url.ParseQuery(u.RawQuery))
}
