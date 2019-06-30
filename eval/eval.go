package eval

import (
	"monkey-go/ast"
	"monkey-go/object"
)

func Eval(node ast.Node) object.Object {

	switch node := node.(type) {

	case *ast.Program:
		return evalStatements(node.Statements)

	case *ast.ExpressionStatement:
		return Eval(node.Expression)

	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}

	case *ast.Boolean:
		return boolean(node.Value)

	}

	return nil
}

func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range stmts {
		result = Eval(statement)
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
