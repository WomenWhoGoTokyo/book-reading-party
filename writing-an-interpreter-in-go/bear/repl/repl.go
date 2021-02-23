package repl

import (
	"bufio"
	"fmt"
	"io"

	"bear/lexer"
	"bear/parser"
)

const PROMPT = "ʕ•ᴥ•ʔ>> "

const BEAR_FACE = `
|:::::::::::::::::::::::::::::::
|"￣￣ﾞﾞヾつ::::::::::::::::::::::
|,ノ　　ヽ,ヽ:::::::::::::::::::::::::
| 0　　　0 i'ﾞ￣￣ﾞﾞﾞ'':::::::::::::::
|　(_o_)　　ﾐノ　　ヽヾつ::::::::::
|　ヽノ　　ノ 0　　 0  i::::::::::
{ヽ,__　　 )   (_o_)　',ﾐ:::::::
| ヽ　　　/　 ヽノ　 ,ノ::::::
`

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, BEAR_FACE)
	io.WriteString(out, "Woops! Web ran into some bear business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
