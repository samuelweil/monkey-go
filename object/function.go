package object

import (
	"bytes"
	"strings"

	"monkey-go/ast"
	"monkey-go/token"
)

type Function struct {
	Token      token.Token
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (fl *Function) Type() Type { return FUNCTION }
func (fl *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())

	return out.String()
}

func (fl *Function) Truthy() bool { return true }
