package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

var data = [][]string{
	{"name", "price"},
	{"りんご", "250"},
	{"どんぐり", "100"},
	{"ビール", "500"},
}

func main() {
	file, err := os.Create("sample.csv")
	if err != nil {
		fmt.Println(err)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	writer := csv.NewWriter(io.MultiWriter(file, os.Stdout))

	for _, d := range data {
		if err := writer.Write(d); err != nil {
			fmt.Println(err)
		}
	}
	writer.Flush()

	if err := writer.Error(); err != nil {
		fmt.Println(err)
	}
}
