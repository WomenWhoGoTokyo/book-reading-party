package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func CopyN(dest io.Writer, src io.Reader, length int64) (written int64, err error) {
	if length <= 0 {
		return length, nil
	}
	buf := make([]byte, length)
	written, err = io.CopyBuffer(dest, io.LimitReader(src, length), buf)
	if written < length {
		err = io.EOF
	}
	return
}

func main() {
	r := strings.NewReader("some io.Reader stream to be read")

	if _, err := CopyN(os.Stdout, r, 10); err != nil {
		log.Fatal(err)
	}
}
