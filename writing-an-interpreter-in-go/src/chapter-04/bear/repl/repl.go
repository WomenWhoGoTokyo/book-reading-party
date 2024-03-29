package repl

import (
	"bufio"
	"fmt"
	"io"

	"bear/evaluator"
	"bear/lexer"
	"bear/object"
	"bear/parser"
)

const PROMPT = "ʕ ꉺᴥ ꉺʔ>> "

const BEAR_FACE = `
|:::::::::::::::::::::::::::::::
|"￣￣ﾞﾞヾつ::::::::::::::::::::::
|,ノ　　ヽ,ヽ:::::::::::::::::::::::::
| ꄱ　  ꄱ  i'ﾞ￣￣ﾞﾞﾞ'':::::::::::::::
|　(_o_)　　ﾐノ　　ヽヾつ:::::::::::
|  ヽノ　　ノ ꉺ　  ꉺ  i::::::::::
{ヽ,__　　 )   (_o_)　',ﾐ::::::::
| ヽ　　　/    ヽノ ,ノ::::🍯🍯🍯
`

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()
	macroEnv := object.NewEnvironment()

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

		evaluator.DefineMacros(program, macroEnv)
		expanded := evaluator.ExpandMacros(program, macroEnv)

		evaluated := evaluator.Eval(expanded, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
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
