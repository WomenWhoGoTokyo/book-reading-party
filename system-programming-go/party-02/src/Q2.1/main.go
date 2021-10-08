package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("q2.1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	writer := io.Writer(file)

	if _, err = fmt.Fprintf(writer,
		"%s のおこづかいは %d 円です。%d 円の買い物をしました。おこづかいが %f になりました。",
		"Gopher", 500, 200, 0.4,
	); err != nil {
		fmt.Println(err)
	}
}
