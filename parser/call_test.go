package parser

import (
	"monkey-go/ast"
	"testing"

	"github.com/samuelweil/go-tools/testing/assert"
)

func TestCallExpressions(t *testing.T) {
	assert := assert.New(t)

	input := `add(1, 2 * 3, 4 + 5)`

	p := New(input)
	prgm := p.ParseProgram()
	checkParserErrors(t, p)

	assert.Eq(len(prgm.Statements), 1)

	stmt, ok := prgm.Statements[0].(*ast.ExpressionStatement)
	assert.True(ok, "stmt is not *ast.ExpressionStatement. Got %T", prgm.Statements[0])

	exp, ok := stmt.Expression.(*ast.CallExpression)
	assert.True(ok, "exp is not *ast.CallExpression. Got %T", stmt.Expression)

	assert.NoError(testIdentifier(exp.Function, "add"))

	assert.Eq(len(exp.Arguments), 3)

	assert.NoError(testLiteralExpression(exp.Arguments[0], 1))
	assert.NoError(testInfixExpression(exp.Arguments[1], infixTest{2, "*", 3}))
	assert.NoError(testInfixExpression(exp.Arguments[2], infixTest{4, "+", 5}))
}
