package parser

import (
	"monkey-go/ast"
	"testing"

	"github.com/samuelweil/go-tools/testing/assert"
)

func TestParsingIndexExpression(t *testing.T) {

	assert := assert.New(t)

	input := "myArray[1 + 1]"

	p := New(input)
	prgm := p.ParseProgram()
	checkParserErrors(t, p)

	stmt, _ := prgm.Statements[0].(*ast.ExpressionStatement)
	indexExp, ok := stmt.Expression.(*ast.IndexExpression)

	assert.True(ok, "exp not *ast.IndexExpression. Got %T", stmt.Expression)
	assert.NoError(testIdentifier(indexExp.Left, "myArray"))
	assert.NoError(testInfixExpression(indexExp.Index, infixTest{1, "+", 1}))
}
