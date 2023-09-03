package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		contents := forecast()
		w.Write([]byte(contents))
	})
	http.ListenAndServe(":8080", nil)
}

func forecast() string {
	return "xxxxxxxxxxxxxxxx"
}
