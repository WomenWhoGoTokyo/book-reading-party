package main

import (
	"crypto/rand"
	"io"
	"os"
)

// mainはランダムな内容で新しいファイルを作成する。
func main() {
	file, err := os.Create("./random")
	if err != nil {
		panic(err)
	}

	io.CopyN(file, rand.Reader, 1024)
}
