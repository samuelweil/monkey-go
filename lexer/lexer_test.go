package lexer

import (
	"monkey-go/token"
	"testing"

	"github.com/samuelweil/go-tools/testing/assert"
)

func validateLexer(t *testing.T, inp string, exp []token.Token) {

	assert := assert.New(t)

	lexer := New(inp)

	for _, tt := range exp {
		tok := lexer.NextToken()

		assert.Eq(tok, tt)
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

func TestMultiCharOperators(t *testing.T) {
	input := `== != <= >= && ||`

	tests := []token.Token{
		token.Eq(),
		token.NotEq(),
		token.LessEq(),
		token.GreatEq(),
		token.And(),
		token.Or(),
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

func TestString(t *testing.T) {
	input := `
	"Hello, World"
	'Hello, World'
	"Other string"
	"This is Sam's test"
	'Test double quote "'
	`

	tests := []token.Token{
		token.String("Hello, World"),
		token.String("Hello, World"),
		token.String("Other string"),
		token.String("This is Sam's test"),
		token.String("Test double quote \""),
	}

	validateLexer(t, input, tests)
}

func TestArray(t *testing.T) {
	input := `
	[1, 2, 3]
	let x = ["xyz", "abc"]
	x[0]
	x[y]`

	tests := []token.Token{
		token.LBrack(), token.Int("1"), token.Comma(), token.Int("2"), token.Comma(), token.Int("3"), token.RBrack(),
		token.Let(), token.Ident("x"), token.Assign(), token.LBrack(), token.String("xyz"), token.Comma(), token.String("abc"), token.RBrack(),
		token.Ident("x"), token.LBrack(), token.Int("0"), token.RBrack(),
		token.Ident("x"), token.LBrack(), token.Ident("y"), token.RBrack(),
	}

	validateLexer(t, input, tests)
}

func TestLexer(t *testing.T) {
	input := `
	let five = 5;	
	let ten = 10;

	let hello = "world";
	let hello = 'world';

	let add = fn(x, y) {
		x + y;
	};

	let result = add(five, ten);

	if 5 != 10 {
		ret true;
	} else {
		ret false;
	}

	for {

	}

	`
	tests := []token.Token{
		token.Let(), token.Ident("five"), token.Assign(), token.Int("5"), token.SemiColon(),
		token.Let(), token.Ident("ten"), token.Assign(), token.Int("10"), token.SemiColon(),
		token.Let(), token.Ident("hello"), token.Assign(), token.String("world"), token.SemiColon(),
		token.Let(), token.Ident("hello"), token.Assign(), token.String("world"), token.SemiColon(),
		token.Let(), token.Ident("add"), token.Assign(), token.Function(), token.LParen(),
		token.Ident("x"), token.Comma(), token.Ident("y"), token.RParen(), token.LBrace(),
		token.Ident("x"), token.Plus(), token.Ident("y"), token.SemiColon(),
		token.RBrace(), token.SemiColon(),
		token.Let(), token.Ident("result"), token.Assign(), token.Ident("add"), token.LParen(),
		token.Ident("five"), token.Comma(), token.Ident("ten"), token.RParen(), token.SemiColon(),
		token.If(), token.Int("5"), token.NotEq(), token.Int("10"), token.LBrace(),
		token.Return(), token.True(), token.SemiColon(),
		token.RBrace(), token.Else(), token.LBrace(),
		token.Return(), token.False(), token.SemiColon(),
		token.RBrace(),
		token.For(), token.LBrace(),
		token.RBrace(),
		token.Eof(),
	}

	validateLexer(t, input, tests)
}
