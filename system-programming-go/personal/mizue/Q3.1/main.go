package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Arg1: Old file path\nArg2: New file path")
		os.Exit(1)
	}

	oldFilePath := os.Args[1]
	newFilePath := os.Args[2]

	oldFile, err := os.Open(oldFilePath)
	if err != nil {
		panic(err)
	}
	defer oldFile.Close()

	newFile, err := os.Create(newFilePath)
	if err != nil {
		panic(nil)
	}
	defer newFile.Close()

	_, err = io.Copy(newFile, oldFile)
	if err != nil {
		panic(err)
	}
}
