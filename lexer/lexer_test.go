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

func TestOneCharTokens(t *testing.T) {
	input := `=+(){}[],;`

	tests := []token.Token{
		token.Assign(),
		token.Plus(),
		token.LParen(),
		token.RParen(),
		token.LBrace(),
		token.RBrace(),
		token.LBrack(),
		token.RBrack(),
		token.Comma(),
		token.SemiColon(),
	}

	lexer := New(input)

	for _, tt := range tests {
		tok := lexer.NextToken()

		assertEq(t, tt, tok)
	}
}

func TestLexer(t *testing.T) {
	input := `
	let five = 5;	
	let ten = 10;

	let add = fn(x, y) {
		x + y;
	};

	let result = add(five, ten);
	`
	tests := []token.Token{
		token.Let(), token.Ident("five"), token.Assign(), token.Int("5"), token.SemiColon(),
		token.Let(), token.Ident("ten"), token.Assign(), token.Int("10"), token.SemiColon(),
		token.Let(), token.Ident("add"), token.Assign(), token.Function(), token.LParen(),
		token.Ident("x"), token.Comma(), token.Ident("y"), token.RParen(), token.LBrace(),
		token.Ident("x"), token.Plus(), token.Ident("y"), token.SemiColon(),
		token.RBrace(), token.SemiColon(),
		token.Let(), token.Ident("result"), token.Assign(), token.Ident("add"), token.LParen(),
		token.Ident("five"), token.Comma(), token.Ident("ten"), token.RParen(), token.SemiColon(),
		token.Eof(),
	}

	lexer := New(input)

	for _, tt := range tests {
		tok := lexer.NextToken()

		assertEq(t, tt, tok)
	}

}
