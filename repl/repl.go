package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey-go/eval"
	"monkey-go/object"
	"monkey-go/parser"
)

const PROMPT = ">> "

type Options struct {
	Debug bool
}

func Start(in io.Reader, out io.Writer, opts Options) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		p := parser.New(line)

		prgm := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		if evaluated := eval.Eval(prgm, env); evaluated != nil {

			if opts.Debug {
				fmt.Printf("%T (%+v)", evaluated, evaluated)
			}

			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		} else {
			break
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, err := range errors {
		io.WriteString(out, fmt.Sprintf("\t%s\n", err))
	}
}
