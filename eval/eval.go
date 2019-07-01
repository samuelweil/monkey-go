package eval

import (
	"fmt"
	"monkey-go/ast"
	"monkey-go/object"
)

func Eval(node ast.Node) object.Object {

	switch node := node.(type) {

	case *ast.Program:
		return evalProgram(node)

	case *ast.ExpressionStatement:
		return Eval(node.Expression)

	case *ast.BlockStatement:
		return evalBlockStatement(node)

	case *ast.ReturnStatement:
		val := Eval(node.ReturnValue)
		return &object.ReturnValue{Value: val}

	case *ast.PrefixExpression:
		right := Eval(node.Right)
		return evalPrefixExpression(node.Operator, right)

	case *ast.InfixExpression:
		left := Eval(node.Left)
		right := Eval(node.Right)
		return evalInfixExpression(node.Operator, left, right)

	case *ast.IfExpression:
		return evalIfExpression(node)

	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}

	case *ast.Boolean:
		return boolean(node.Value)

	}

	return Null
}

func evalProgram(program *ast.Program) object.Object {
	var result object.Object

	for _, statement := range program.Statements {
		result = Eval(statement)

		switch result := result.(type) {
		case *object.ReturnValue:
			return result.Value
		case *object.Error:
			return result
		}
	}

	return result
}

func evalBlockStatement(bs *ast.BlockStatement) object.Object {
	var result object.Object

	for _, statement := range bs.Statements {
		result = Eval(statement)

		if result != nil {

			typ := result.Type()
			if typ == object.RETURN_VALUE || typ == object.ERROR {
				return result
			}
		}
	}

	return result
}

var (
	True  = &object.Boolean{Value: true}
	False = &object.Boolean{Value: false}
	Null  = &object.Null{}
)

func boolean(b bool) *object.Boolean {
	if b {
		return True
	}

	return False
}

func evalPrefixExpression(op string, obj object.Object) object.Object {

	if evaluator, ok := prefixEvals[op]; ok {
		return evaluator(obj)
	}

	return newError("unknown operator: %s%s", op, obj.Type())
}

func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{
		Message: fmt.Sprintf(format, a...),
	}
}
