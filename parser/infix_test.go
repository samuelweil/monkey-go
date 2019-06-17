package parser

import (
	"monkey-go/ast"
	"testing"

	"github.com/samuelweil/go-tools/testing/assert"
)

func TestParsingInfixExpressions(t *testing.T) {
	assert := assert.New(t)

	infixTests := []struct {
		input      string
		leftValue  int
		operator   string
		rightValue int
	}{
		{"5 + 5", 5, "+", 5},
		{"5 - 5", 5, "-", 5},
		{"5 == 5", 5, "==", 5},
	}

	for _, tt := range infixTests {
		p := New(tt.input)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		assert.Eq(len(program.Statements), 1)

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		assert.True(ok, "%s is not an *ast.ExpressionStatement", program.Statements[0])

		exp, ok := stmt.Expression.(*ast.InfixExpression)
		assert.True(ok, "%s is not an *ast.ExpressionStatement", exp)

		if !testIntegerLiteral(t, exp.Left, tt.leftValue) {
			return
		}

		assert.Eq(exp.Operator, tt.operator)

		testIntegerLiteral(t, exp.Right, tt.rightValue)
	}
}
