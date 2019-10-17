package main

import (
	"os"

	"github.com/BigOyayubi/gosimplehttpd"
)

func main() {
	os.Exit((&gosimplehttpd.OptParse{ErrStream: os.Stderr, OutStream: os.Stdout}).Run(os.Args[1:]))
}
