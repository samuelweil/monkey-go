package parser

import (
	"monkey-go/ast"
	"testing"

	"github.com/samuelweil/go-tools/testing/assert"
)

func TestIfExpression(t *testing.T) {
	assert := assert.New(t)

	input := `if (x < y) { x }`

	p := New(input)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	assert.Eq(len(program.Statements), 1)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(ok, "stmt not *ast.ExpressionStatement. Got %T", stmt)

	ifExpr, ok := stmt.Expression.(*ast.IfExpression)
	assert.True(ok, "ifExpr not *ast.IfExpression. Got %T", ifExpr)

	assert.NoError(testInfixExpression(ifExpr.Condition, infixTest{"x", "<", "y"}))

	assert.Eq(len(ifExpr.Consequence.Statements), 1)

	consequence, ok := ifExpr.Consequence.Statements[0].(*ast.ExpressionStatement)
	assert.True(ok, "consequence is not *ast.ExpressionStatment. Got %T", consequence)

	assert.NoError(testIdentifier(consequence.Expression, "x"))

	assert.Nil(ifExpr.Alternative)
}

func TestIfElseExpression(t *testing.T) {
	assert := assert.New(t)

	input := `if (x < y) { x } else { y }`

	p := New(input)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	assert.Eq(len(program.Statements), 1)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(ok, "stmt not *ast.ExpressionStatement. Got %T", stmt)

	ifExpr, ok := stmt.Expression.(*ast.IfExpression)
	assert.True(ok, "ifExpr not *ast.IfExpression. Got %T", ifExpr)

	assert.NoError(testInfixExpression(ifExpr.Condition, infixTest{"x", "<", "y"}))

	assert.Eq(len(ifExpr.Consequence.Statements), 1)

	consequence, ok := ifExpr.Consequence.Statements[0].(*ast.ExpressionStatement)
	assert.True(ok, "consequence is not *ast.ExpressionStatment. Got %T", consequence)

	assert.NoError(testIdentifier(consequence.Expression, "x"))

	assert.Eq(len(ifExpr.Alternative.Statements), 1)

	alt, ok := ifExpr.Alternative.Statements[0].(*ast.ExpressionStatement)
	assert.True(ok, "alternative is not *ast.ExpressionStatement. Got %T", alt)

	assert.NoError(testIdentifier(alt.Expression, "y"))
}
