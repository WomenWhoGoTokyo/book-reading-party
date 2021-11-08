package main

import (
	"io"
	"os"
	"strings"
)

// mainは独自実装したmyCopyNを実際に呼び出す。
func main() {
	myCopyN(os.Stdout, strings.NewReader("01234567890123456789012345678901234567890123456789"), 3)
}

// myCopyNはio.CopyN(dest io.Writer, src io.Reader, length int)を独自実装した処理。
func myCopyN(dest io.Writer, src io.Reader, length int) {
	// TODO: 結局io.CopyNの中身と一緒のことをしており、他の方法がないか再検討する。
	io.Copy(dest, io.LimitReader(src, int64(length)))
}
