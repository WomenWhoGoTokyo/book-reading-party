package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var data = map[string]string{
	"name":   "Gopher",
	"strong": "fast",
	"hobby":  "Go",
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")

	writer := gzip.NewWriter(w)
	defer func() {
		err := writer.Close()
		if err != nil {
			http.Error(w, fmt.Sprintf("gzip error: %v", err), http.StatusInternalServerError)
		}
	}()

	encoder := json.NewEncoder(io.MultiWriter(writer, os.Stdout))
	encoder.SetIndent("", "    ")
	if err := encoder.Encode(data); err != nil {
		http.Error(w, fmt.Sprintf("encode error: %v", err), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
