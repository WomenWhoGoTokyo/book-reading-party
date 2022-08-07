package main

import (
	"fmt"
	"os"
	"strings"
)

func readFile(path string) chan string {
	promise := make(chan string)
	go func() {
		content, err := os.ReadFile(path)
		if err != nil {
			fmt.Printf("readerr %s\n", err.Error())
			close(promise)
		} else {
			promise <- string(content)
		}
	}()
	return promise
}

func printFunc(futureSource chan string) chan []string {
	promise := make(chan []string)
	go func() {
		var result []string
		for _, line := range strings.Split(<-futureSource, "\n") {
			if strings.HasPrefix(line, "func ") {
				result = append(result, line)
			}
		}
		promise <- result
	}()
	return promise
}

func main() {
	futureSource := readFile("future_promise.go")
	futureFuncs := printFunc(futureSource)
	fmt.Println(strings.Join(<-futureFuncs, "\n"))
}
