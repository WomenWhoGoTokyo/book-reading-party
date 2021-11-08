package main

import (
	"io"
	"os"
)

// mainは古いファイルの中身をコピーし、新しく作成するファイルに書き込む。
func main() {
	oldFile, err := os.Open("./old.txt")
	if err != nil {
		panic(err)
	}
	defer oldFile.Close()

	newFile, err := os.Create("./new.txt")
	if err != nil {
		panic(err)
	}

	io.Copy(newFile, oldFile)
}
