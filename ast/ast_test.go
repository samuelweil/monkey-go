package ast

import (
	"monkey-go/token"
	"testing"
	"tools/testing/assert"
)

func TestString(t *testing.T) {
	assert := assert.New(t)
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Let(),
				Name: &Identifier{
					Token: token.Ident("myVar"),
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Ident("anotherVar"),
					Value: "anotherVar",
				},
			},
		},
	}

	assert.Eq(program.String(), "let myVar = anotherVar;")
}
