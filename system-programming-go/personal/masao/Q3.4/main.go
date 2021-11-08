package main

import (
	"io"
	"net/http"
	"os"
)

// mainはwebサーバーの起動を行う。
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

// handlerはアクセスが来たらzipファイルを返す処理。
func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=result.zip")

	zipFile, err := os.Open("./q4.zip")
	if err != nil {
		panic(err)
	}
	defer zipFile.Close()

	io.Copy(w, zipFile)
}
