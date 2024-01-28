package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var c chan int

func handle(int) {}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Arg1: number of seconds")
		os.Exit(1)
	}

	arg1, _ := strconv.ParseInt(os.Args[1], 10, 64)
	sec := time.Duration(arg1)

	select {
	case m := <-c:
		handle(m)
	case <-time.After(sec * time.Second):
		fmt.Println("timed out")
	}
}
