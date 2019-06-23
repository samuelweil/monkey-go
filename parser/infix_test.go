package parser

import (
	"fmt"
	"monkey-go/ast"
	"testing"

	"github.com/samuelweil/go-tools/testing/assert"
	"github.com/samuelweil/go-tools/testing/check"
)

func TestParsingInfixExpressions(t *testing.T) {
	assert := assert.New(t)
	check := check.New(t)

	infixTests := []struct {
		input    string
		expected infixTest
	}{
		{"5 + 5", infixTest{5, "+", 5}},
		{"5 - 5", infixTest{5, "-", 5}},
		{"5 == 5", infixTest{5, "==", 5}},
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

		check.NoError(testInfixExpression(exp, tt.expected))
	}
}

type infixTest struct {
	Left     interface{}
	Operator string
	Right    interface{}
}

func testInfixExpression(exp ast.Expression, expected infixTest) error {

	opExp, ok := exp.(*ast.InfixExpression)
	if !ok {
		return fmt.Errorf("exp is not *ast.Expression. Got %T", exp)
	}

	if e := testLiteralExpression(opExp.Left, expected.Left); e != nil {
		return e
	}

	if opExp.Operator != expected.Operator {
		return fmt.Errorf("exp.Operator is not %s. Got %s", expected.Operator, opExp.Operator)
	}

	if e := testLiteralExpression(opExp.Right, expected.Right); e != nil {
		return e
	}

	return nil
}
