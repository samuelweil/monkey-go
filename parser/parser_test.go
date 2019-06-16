package parser

import (
	"monkey-go/ast"
	"testing"
	"tools/testing/assert"
	"tools/testing/check"
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

func TestReturnStatements(t *testing.T) {
	assert := assert.New(t)
	check := check.New(t)

	input := `
	ret 5;
	ret 10;
	ret 993322;
	`
	p := New(input)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	assert.Eq(len(program.Statements), 3)

	for _, stmt := range program.Statements {
		retStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement, got=%T", stmt)
			continue
		}

		check.Eq(retStmt.TokenLiteral(), "ret")
	}
}

func TestIdentifierExpression(t *testing.T) {
	assert := assert.New(t)
	check := check.New(t)

	input := "foobar;"

	p := New(input)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	assert.Eq(len(program.Statements), 1)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(ok, "%s is not an *ast.ExpressionStatement", stmt)

	ident, ok := stmt.Expression.(*ast.Identifier)
	assert.True(ok, "%s is not an *ast.Identifier", ident)

	check.Eq(ident.Value, "foobar")
	check.Eq(ident.TokenLiteral(), "foobar")
}
