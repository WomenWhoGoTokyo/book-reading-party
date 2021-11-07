package main

import (
	"fmt"
	"strings"
)

var src1 = "123 1.234 1.0e4 test"
var src2 = `123
1.234
1.0e4
test
`

func main() {
	var i int
	var f, g float64
	var s string

	reader := strings.NewReader(src1)
	fmt.Fscan(reader, &i, &f, &g, &s)
	fmt.Printf("i=%#v f=%#v g=%#v s=%#v\n", i, f, g, s)

	reader = strings.NewReader(src2)
	fmt.Fscanln(reader, &i, &f, &g, &s)
	fmt.Printf("i=%#v f=%#v g=%#v s=%#v\n", i, f, g, s)
}
