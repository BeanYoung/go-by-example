package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	buildCmd := flag.NewFlagSet("build", flag.ExitOnError)
	nameFlag := buildCmd.String("name", "", "name")

	if len(os.Args) < 2 {
		panic("sub command is missing")
	}

	switch os.Args[1] {
	case "build":
		buildCmd.Parse(os.Args[2:])
		fmt.Println(*nameFlag)
	default:
		panic("unknown subcommand")
	}
}
