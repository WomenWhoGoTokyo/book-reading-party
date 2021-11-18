package main

import (
	"io"
	"os"
	"strings"
)

var (
	computer    = strings.NewReader("COMPUTER")
	system      = strings.NewReader("SYSTEM")
	programming = strings.NewReader("PROGRAMMING")
)

func main() {

	var stream io.Reader

	r, w := io.Pipe()
	go func() {
		io.Copy(w, io.NewSectionReader(programming, 5, 1))
		io.Copy(w, io.LimitReader(system, 1))
		io.Copy(w, io.LimitReader(computer, 1))
		io.Copy(w, io.NewSectionReader(programming, 8, 1))
		io.Copy(w, io.NewSectionReader(programming, 8, 1))
		w.Close()
	}()

	stream = r
	// ASCII
	io.Copy(os.Stdout, stream)

}
