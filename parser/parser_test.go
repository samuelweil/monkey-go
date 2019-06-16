package parser

import (
	"monkey-go/ast"
	"monkey-go/testing/assert"
	"monkey-go/testing/check"
	"testing"
)

func TestLetStatement(t *testing.T) {

	assert := assert.New(t)

	input := `
	let x = 5;
	let y = 10;
	let foobar = 838383;
	`
	p := New(input)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	assert.Eq(len(program.Statements), 3)

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
	check := check.New(t)

	stmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("Expected *ast.LetStatement, Got %T", s)
		return false
	}

	if !check.Eq(stmt.Name.Value, name) {
		return false
	}

	if !check.Eq(stmt.Name.TokenLiteral(), name) {
		return false
	}

	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("%d parsing errors", len(errors))

	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
