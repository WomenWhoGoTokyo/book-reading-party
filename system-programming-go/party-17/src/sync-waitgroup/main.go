package main

import (
	"fmt"
	"sync"
)

var id int

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("仕事1")
		wg.Done()
	}()

	go func() {
		fmt.Println("仕事2")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("終了")
}
