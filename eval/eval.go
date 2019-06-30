package eval

import (
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

	return nil
}

func evalProgram(program *ast.Program) object.Object {
	var result object.Object

	for _, statement := range program.Statements {
		result = Eval(statement)

		if returnValue, ok := result.(*object.ReturnValue); ok {
			return returnValue.Value
		}
	}

	return result
}

func evalBlockStatement(bs *ast.BlockStatement) object.Object {
	var result object.Object

	for _, statement := range bs.Statements {
		result = Eval(statement)

		if result != nil && result.Type() == object.RETURN_VALUE {
			return result
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

	return nil
}
