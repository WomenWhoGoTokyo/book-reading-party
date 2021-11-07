package main

import (
	"bufio"
	"fmt"
	"strings"
)

var source2 = `1行目
2行目
3行目`

func main() {
	scanner := bufio.NewScanner(strings.NewReader(source2))
	for scanner.Scan() {
		fmt.Printf("%#v\n", scanner.Text())
	}
}
