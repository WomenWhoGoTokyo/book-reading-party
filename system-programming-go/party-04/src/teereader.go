package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
)

func main() {
	var buffer bytes.Buffer
	reader := bytes.NewBufferString("Example of io.TeeReader\n")
	teeReader := io.TeeReader(reader, &buffer)
	// 読み捨てる
	_, _ = ioutil.ReadAll(teeReader)
	// バッファに残っている
	fmt.Println(buffer.String())
}
