package main

import (
	"fmt"
	"io"
	"os"

	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"
)

func main() {
	var out io.Writer
	if isatty.IsTerminal(os.Stdout.Fd()) {
		out = colorable.NewColorableStderr()
	} else {
		out = colorable.NewNonColorable(os.Stdout)
	}
	if isatty.IsTerminal(os.Stdin.Fd()) {
		fmt.Fprintln(out, "stdin: terminal")
	} else {
		fmt.Println("stdin: pipe")
	}
	if isatty.IsTerminal(os.Stdout.Fd()) {
		fmt.Fprintln(out, "stdout: terminal")
	} else {
		fmt.Println("stdout: pipe")
	}
	if isatty.IsTerminal(os.Stderr.Fd()) {
		fmt.Fprintln(out, "stderr: terminal")
	} else {
		fmt.Println("stderr: pipe")
	}
}
