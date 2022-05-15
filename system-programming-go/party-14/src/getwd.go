package main

import (
	"fmt"
	"os"
)

func main() {
	wd, _ := os.Gewd()
	fmt.Println(wd)
}
