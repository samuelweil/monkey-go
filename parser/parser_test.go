package parser

import (
	"fmt"
	"monkey-go/ast"
	"testing"

	"github.com/samuelweil/go-tools/testing/assert"
	"github.com/samuelweil/go-tools/testing/check"
)

func TestLetStatement(t *testing.T) {

	assert := assert.New(t)
	check := check.New(t)

	tests := []struct {
		input              string
		expectedIdentifier string
		expectedValue      interface{}
	}{
		{"let x = 5;", "x", 5},
		{"let y = true", "y", true},
		{"let foobar = y", "foobar", "y"},
	}

	for _, tt := range tests {

		p := New(tt.input)
		prgm := p.ParseProgram()
		checkParserErrors(t, p)

		assert.Eq(len(prgm.Statements), 1)

		stmt := prgm.Statements[0]
		assert.NoError(testLetStatement(stmt, tt.expectedIdentifier))

		val := stmt.(*ast.LetStatement).Value
		check.NoError(testLiteralExpression(val, tt.expectedValue))
	}
}

func testLetStatement(s ast.Statement, name string) error {
	stmt, ok := s.(*ast.LetStatement)
	if !ok {
		return fmt.Errorf("Expected *ast.LetStatement, Got %T", s)
	}

	return testIdentifier(stmt.Name, name)
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

	tests := []struct {
		input         string
		expectedValue interface{}
	}{
		{"ret 5;", 5},
		{"ret 10;", 10},
		{"ret false", false},
		{"ret x;", "x"},
	}

	for _, tt := range tests {
		p := New(tt.input)
		prgm := p.ParseProgram()
		checkParserErrors(t, p)

		stmt, ok := prgm.Statements[0].(*ast.ReturnStatement)
		assert.True(ok, "stmt is not a *ast.ReturnStatement. Got %T", prgm.Statements[0])

		check.NoError(testLiteralExpression(stmt.ReturnValue, tt.expectedValue))
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
	assert.True(ok, "%s is not an *ast.ExpressionStatement", program.Statements[0])

	ident, ok := stmt.Expression.(*ast.Identifier)
	assert.True(ok, "%s is not an *ast.Identifier", stmt.Expression)

	check.Eq(ident.Value, "foobar")
	check.Eq(ident.TokenLiteral(), "foobar")
}

func TestIntegerLiteralExpression(t *testing.T) {
	assert := assert.New(t)

	input := "5;"

	p := New(input)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	assert.Eq(len(program.Statements), 1)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(ok, "%s is not a *ast.ExpressionStatement", program.Statements[0])

	assert.NoError(testIntegerLiteral(stmt.Expression, 5))
}

func TestStringLiteralExpression(t *testing.T) {
	assert := assert.New(t)

	input := `"hello, world"`

	p := New(input)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	assert.Eq(len(program.Statements), 1)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(ok, "%s is not a *ast.ExpressionStatement", program.Statements[0])

	assert.NoError(testStringLiteral(stmt.Expression, "hello, world"))

}

func TestParsingPrefixExpressions(t *testing.T) {
	assert := assert.New(t)
	check := check.New(t)

	prefixTests := []struct {
		input    string
		operator string
		value    interface{}
	}{
		{"!5;", "!", 5},
		{"-15;", "-", 15},
		{"!true", "!", true},
		{"!false", "!", false},
	}

	for _, tt := range prefixTests {
		p := New(tt.input)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		assert.Eq(len(program.Statements), 1)

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		assert.True(ok, "%s is not an *ast.ExpressionStatement", program.Statements[0])

		e := testParsingPrefixExpressions(stmt.Expression, tt.operator, tt.value)
		check.NoError(e)
	}
}

func testParsingPrefixExpressions(exp ast.Expression, operator string, value interface{}) error {

	preExp, ok := exp.(*ast.PrefixExpression)
	if !ok {
		return fmt.Errorf("exp is not *ast.PrefixExpression. Got %T", exp)
	}

	if preExp.Operator != operator {
		return fmt.Errorf("exp.Operator is not %s. Got %s", preExp.Operator, operator)
	}

	return testLiteralExpression(preExp.Right, value)
}

func TestParsingBoolean(t *testing.T) {

	assert := assert.New(t)
	check := check.New(t)

	boolTests := []struct {
		input string
		value bool
	}{
		{"true", true},
		{"false", false},
	}

	for _, tt := range boolTests {
		p := New(tt.input)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		assert.Eq(len(program.Statements), 1)

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		assert.True(ok, "%s is not an *ast.ExpressionStatement", program.Statements[0])

		check.NoError(testBoolean(stmt.Expression, tt.value))
	}
}

func testIdentifier(exp ast.Expression, value string) error {

	ident, ok := exp.(*ast.Identifier)
	if !ok {
		return fmt.Errorf("exp not *ast.Identifier. Got %T", exp)
	}

	if ident.Value != value {
		return fmt.Errorf("ident.Value not %s. Got %s", value, ident.Value)
	}

	if ident.TokenLiteral() != value {
		return fmt.Errorf("ident.TokenLiteral() not %s. Got %s", value, ident.TokenLiteral())
	}

	return nil
}

func testLiteralExpression(exp ast.Expression, expected interface{}) error {

	switch v := expected.(type) {
	case int:
		return testIntegerLiteral(exp, v)
	case string:
		return testIdentifier(exp, v)
	case bool:
		return testBoolean(exp, v)
	default:
		return fmt.Errorf("Unhandled expression type %T", exp)
	}
}

func testIntegerLiteral(il ast.Expression, value int) error {

	literal, ok := il.(*ast.IntegerLiteral)
	if !ok {
		return fmt.Errorf("%s is not a *ast.IntegerLiteral", il)
	}

	if literal.Value != value {
		return fmt.Errorf("literal.Value is not %v. Got %v", value, literal.Value)
	}

	return nil
}

func testStringLiteral(sl ast.Expression, value string) error {
	literal, ok := sl.(*ast.StringLiteral)

	if !ok {
		return fmt.Errorf("%s is not a *ast.StringLiteral", sl)
	}

	if literal.Value != value {
		return fmt.Errorf("literal.Value is not %v. Got %v", value, literal.Value)
	}

	return nil
}

func testBoolean(exp ast.Expression, value bool) error {

	literal, ok := exp.(*ast.Boolean)
	if !ok {
		return fmt.Errorf("%s is not a *ast.Boolean", exp)
	}

	if literal.Value != value {
		return fmt.Errorf("literal.Value is not %v. Got %v", value, literal.Value)
	}

	if literal.TokenLiteral() != fmt.Sprintf("%t", value) {
		return fmt.Errorf("literal.TokenLiteral() is not %t. Got %s", value, literal.TokenLiteral())
	}

	return nil
}
