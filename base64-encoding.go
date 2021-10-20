package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	s := "http://www.baidu.com/?a=b&c=d"

	fmt.Println(s)
	fmt.Println(base64.StdEncoding.EncodeToString([]byte(s)))
	fmt.Println(base64.URLEncoding.EncodeToString([]byte(s)))
}
