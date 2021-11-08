package main

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

// mainは文字列strings.Readerを使ってzipファイルを作成する。
func main() {
	// zipファイルの用意
	zipFile, err := os.Create("./result.zip")
	if err != nil {
		panic(err)
	}

	// zipファイルへの書き込み用Writerの用意
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// zipファイル内に1.txtというファイルを追加
	str1Writer, err := zipWriter.Create("result/1.txt")
	if err != nil {
		panic(err)
	}

	// 1.txtに書き込み
	io.Copy(str1Writer, strings.NewReader("Hello!"))
}
