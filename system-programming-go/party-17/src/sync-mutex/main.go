package main

import (
	"fmt"
	"sync"
)

var id int

func generateId(mutex *sync.Mutex) int {
	mutex.Lock()
	id++
	result := id
	mutex.Unlock()
	return result
}

func main() {
	var mutex sync.Mutex
	for i := 0; i < 100; i++ {
		fmt.Println(i)
		go func() {
			fmt.Printf("id %d\n", generateId(&mutex))
		}()
	}
}
