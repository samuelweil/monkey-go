package ast

import (
	"bytes"
	"monkey-go/token"
	"strings"
)

type ArrayLiteral struct {
	Token    token.Token
	Elements []Expression
}

func (al *ArrayLiteral) TokenLiteral() string { return al.Token.Literal }
func (al *ArrayLiteral) String() string {
	var out bytes.Buffer

	elements := []string{}
	for _, el := range al.Elements {
		elements = append(elements, el.String())
	}

	out.WriteString(token.LBRACK)
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString(token.RBRACK)

	return out.String()
}
