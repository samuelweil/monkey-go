package parser

import (
	"monkey-go/ast"
	"testing"

	"github.com/samuelweil/go-tools/testing/assert"
	"github.com/samuelweil/go-tools/testing/check"
)

func TestFunction(t *testing.T) {
	assert := assert.New(t)
	check := check.New(t)

	input := `fn(x, y) { x + y; }`

	p := New(input)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	assert.Eq(len(program.Statements), 1)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(ok, "program.Statements[0] is not *ast.ExpressionStatement. Got %T", program.Statements[0])

	function, ok := stmt.Expression.(*ast.FunctionLiteral)
	assert.True(ok, "stmt.Expression is not a *ast.FunctionLiteral. Got %T", stmt)

	assert.Eq(len(function.Parameters), 2)

	check.NoError(testLiteralExpression(function.Parameters[0], "x"))
	check.NoError(testLiteralExpression(function.Parameters[1], "y"))

	assert.Eq(len(function.Body.Statements), 1)

	bodyStmt, ok := function.Body.Statements[0].(*ast.ExpressionStatement)
	assert.True(ok, "function body is not *ast.ExpressionStatement. Got %T", function.Body.Statements[0])

	check.NoError(testInfixExpression(bodyStmt.Expression, infixTest{"x", "+", "y"}))

}

func TestParseFunctionParameters(t *testing.T) {
	check := check.New(t)

	tests := []struct {
		input     string
		expParams []string
	}{
		{input: "fn() {};", expParams: []string{}},
		{input: "fn(x) {};", expParams: []string{"x"}},
		{input: "fn(x, y, z) {};", expParams: []string{"x", "y", "z"}},
	}

	for _, tt := range tests {
		p := New(tt.input)
		prgm := p.ParseProgram()
		checkParserErrors(t, p)

		stmt := prgm.Statements[0].(*ast.ExpressionStatement)
		fn := stmt.Expression.(*ast.FunctionLiteral)

		check.Eq(len(fn.Parameters), len(tt.expParams))

		for i, ident := range tt.expParams {
			check.NoError(testLiteralExpression(fn.Parameters[i], ident))
		}
	}
}
