package parser

import (
	"monkey-go/ast"
	"testing"

	"github.com/samuelweil/go-tools/testing/assert"
)

func TestParsingArray(t *testing.T) {

	assert := assert.New(t)

	input := "[1, 2 * 2, 3 + 3, 'bob']"

	p := New(input)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt, _ := program.Statements[0].(*ast.ExpressionStatement)
	array, ok := stmt.Expression.(*ast.ArrayLiteral)

	assert.True(ok, "exp not ast.ArrayLiteral. got %T", stmt.Expression)
	assert.Eq(len(array.Elements), 4)

	assert.NoError(testIntegerLiteral(array.Elements[0], 1))
	assert.NoError(testInfixExpression(array.Elements[1], infixTest{2, "*", 2}))
	assert.NoError(testInfixExpression(array.Elements[2], infixTest{3, "+", 3}))
	assert.NoError(testStringLiteral(array.Elements[3], "bob"))
}
