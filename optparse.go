package gosimplehttpd

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const version = "0.0.1"

var revision = "HEAD"

const (
	codeOK = iota
	codeErr
)

// OptParse is struct for argv parse
type OptParse struct {
	OutStream, ErrStream io.Writer
}

// parse argv and start gosimplehttpd
func (optParse *OptParse) Run(argv []string) int {
	// setup flags
	var dir string
	flag.StringVar(&dir, "d", ".", "root directory path for glob")
	var port string
	flag.StringVar(&port, "p", "8080", "port")
	var isHelp bool
	flag.BoolVar(&isHelp, "h", false, "show help")

	flag.Parse()

	if isHelp {
		fmt.Fprintln(optParse.ErrStream, help())
		return codeOK
	}

	var abs, _ = filepath.Abs(dir)

	if _, err := os.Stat(abs); os.IsNotExist(err) {
		fmt.Println("dir exists is not exist")
		return codeErr
	}

	fmt.Println("root dir : ", abs)
	fmt.Println("port     : ", port)

	(&FileHttpHandler{Directory: abs, Port: port}).Awake()

	return codeOK
}

func help() string {
	return fmt.Sprintf(`Usage:
	$ gosimplehttpd [-d="."] [-p=8080]

Utility to emulate asset delivery server
It replys a file to http get request.

Options:
  -d           asset root directory.
  -p           port. > http://localhost:port
  -h           show help
Version: %s`, version)
}
