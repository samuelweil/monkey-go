package lexer

import (
	"monkey-go/token"
	"testing"
)

func assertEq(t *testing.T, exp, got interface{}) {
	if exp != got {
		t.Errorf("Expected %v, got %v", exp, got)
	}
}

func TestNextToken(t *testing.T) {
	input := `=+(){}[],;`

	tests := []token.Token{
		token.FromStr(token.ASSIGN),
		token.FromStr(token.PLUS),
		token.FromStr(token.LPAREN),
		token.FromStr(token.RPAREN),
		token.FromStr(token.LBRACE),
		token.FromStr(token.RBRACE),
		token.FromStr(token.LBRACK),
		token.FromStr(token.RBRACK),
		token.FromStr(token.COMMA),
		token.FromStr(token.SEMICOLON),
	}

	lexer := New(input)

	for _, tt := range tests {
		tok := lexer.NextToken()

		assertEq(t, tt, tok)
	}

}
