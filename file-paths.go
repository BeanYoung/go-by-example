package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	s := "/tmp/123/456"

	fmt.Println(filepath.Join("tmp", "demo", "1"))
	fmt.Println(filepath.IsAbs(s))
	fmt.Println(filepath.Ext("demo.conf"))
	fmt.Println(filepath.Rel("/tmp/123", "/tmp/123/456/789"))
}
