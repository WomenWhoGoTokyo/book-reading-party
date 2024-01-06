package main

import (
	"log"

	"github.com/WomenWhoGoTokyo/proglog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
