package main

import (
	"crypto/rand"
	"os"
)

func main() {
	b := make([]byte, 1024)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}

	newFile, err := os.Create("rand.txt")
	if err != nil {
		panic(nil)
	}
	defer newFile.Close()

	newFile.Write(b)
}
