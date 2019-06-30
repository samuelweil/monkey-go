package eval

import (
	"monkey-go/ast"
	"monkey-go/object"
)

func evalIfExpression(ie *ast.IfExpression) object.Object {
	condition := Eval(ie.Condition)

	if condition.Truthy() {
		return Eval(ie.Consequence)
	} else if ie.Alternative != nil {
		return Eval(ie.Alternative)
	} else {
		return Null
	}
}