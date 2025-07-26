package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, req *http.Request) {
	cnt++
	fmt.Fprintf(w, "hello %d", cnt)
	fmt.Printf("%d %#v\n", cnt, req.UserAgent())
}
func main() {
	fmt.Printf("listen: %s\n", addr)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(addr, nil))
}
