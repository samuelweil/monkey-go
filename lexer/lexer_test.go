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

func validateLexer(t *testing.T, inp string, exp []token.Token) {
	lexer := New(inp)

	for _, tt := range exp {
		tok := lexer.NextToken()

		assertEq(t, tt, tok)
	}
}

func TestOperators(t *testing.T) {
	input := `=+-*/!<>`

	tests := []token.Token{
		token.Assign(),
		token.Plus(),
		token.Minus(),
		token.Asterisk(),
		token.Slash(),
		token.Bang(),
		token.LessThan(),
		token.GreaterThan(),
	}

	validateLexer(t, input, tests)
}

func TestDelimiters(t *testing.T) {
	input := `(){}[],;`

	tests := []token.Token{
		token.LParen(),
		token.RParen(),
		token.LBrace(),
		token.RBrace(),
		token.LBrack(),
		token.RBrack(),
		token.Comma(),
		token.SemiColon(),
	}

	validateLexer(t, input, tests)
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

	validateLexer(t, input, tests)
}
