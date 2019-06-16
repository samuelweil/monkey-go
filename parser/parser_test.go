package parser

import (
	"monkey-go/assert"
	"monkey-go/ast"
	"monkey-go/lexer"
	"testing"
)

func TestLetStatement(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar = 838383;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if n := len(program.Statements); n != 3 {
		t.Fatalf("program has %d statements; expected 3", n)
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]

		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	assert := assert.New(t)

	stmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("Expected *ast.LetStatement, Got %T", s)
		return false
	}

	if !assert.Eq(stmt.Name.Value, name) {
		return false
	}

	if !assert.Eq(stmt.Name.TokenLiteral(), name) {
		return false
	}

	return true
}
