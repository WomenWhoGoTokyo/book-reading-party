package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

var source1 = `1行目
2行目
3行目`

func main() {
	reader := bufio.NewReader(strings.NewReader(source1))
	for {
		line, err := reader.ReadString('\n')
		fmt.Printf("%#v\n", line)
		if err == io.EOF {
			break
		}
	}
}
