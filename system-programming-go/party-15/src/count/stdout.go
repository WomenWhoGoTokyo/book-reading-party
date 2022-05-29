package main

import (
	"bufio"
	"fmt"
	"os/exec"
)

func main() {
	count := exec.Command("./count")
	stdout, _ := count.StdoutPipe()
	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			fmt.Printf("(stdout) %s\n", scanner.Text())
		}
	}()
	err := count.Run()
	if err != nil {
		panic(err)
	}
}
