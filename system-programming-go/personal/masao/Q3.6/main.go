package main

import (
	"io"
	"os"
	"strings"
)

var (
	computer    = strings.NewReader("COMPUTER")
	system      = strings.NewReader("SYSTEM")
	programming = strings.NewReader("PROGRAMMING")
)

// mainは以下の制約のもとASCIIの文字を標準出力に出力する。
// ・使っていいのはioパッケージの内容+基本文法のみ。
// ・文字列リテラルを使用しない。
// ・コメント部以外を変更しない。importするパッケージはstrings, io, osのみ。
func main() {
	var stream io.Reader

	// Aの取得
	stream = io.NewSectionReader(programming, 5, 1)
	io.Copy(os.Stdout, stream)

	// Sの取得
	stream = io.LimitReader(system, 1)
	io.Copy(os.Stdout, stream)

	// Cの取得
	stream = io.LimitReader(computer, 1)
	io.Copy(os.Stdout, stream)

	// Iの取得
	stream = io.NewSectionReader(programming, 8, 1)
	io.Copy(os.Stdout, stream)

	// Iの取得
	stream = io.NewSectionReader(programming, 8, 1)
	io.Copy(os.Stdout, stream)
}
