package main

import (
	"fmt"
	"os/user"
)

func main() {
	my, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Println(my)
}
