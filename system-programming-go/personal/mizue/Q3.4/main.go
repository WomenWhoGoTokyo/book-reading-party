package main

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"strings"
)

var source = "This archive contains some text files."

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=test.zip")

	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()

	f, err := zipWriter.Create("test.txt")
	if err != nil {
		panic(err)
	}

	// 文字列を読み込む
	reader := strings.NewReader(source)

	// zipに書き込む
	io.Copy(f, reader)
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
