package eval

import (
	"fmt"
	"monkey-go/ast"
	"monkey-go/object"
)

func Eval(node ast.Node, env *object.Environment) object.Object {

	switch node := node.(type) {

	case *ast.Program:
		return evalProgram(node, env)

	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)

	case *ast.BlockStatement:
		return evalBlockStatement(node, env)

	case *ast.LetStatement:
		return evalLetStatement(node, env)

	case *ast.ReturnStatement:
		val := Eval(node.ReturnValue, env)
		return &object.ReturnValue{Value: val}
	
	case *ast.FunctionLiteral:
		return &object.Function{ 
			Parameters: node.Parameters,
			Env: env, 
			Body: node.Body,
		}

	case *ast.CallExpression:
		function := Eval(node.Function, env)
		if isError(function) {
			return function
		}

		args := evalExpressions(node.Arguments, env)
		if len(args) == 1 && isError(args[0]) {
			return args[0]
		}

		return applyFunction(function, args)

	case *ast.Identifier:
		return evalIdentifier(node, env)

	case *ast.PrefixExpression:
		right := Eval(node.Right, env)
		return evalPrefixExpression(node.Operator, right, env)

	case *ast.InfixExpression:
		left := Eval(node.Left, env)
		right := Eval(node.Right, env)
		return evalInfixExpression(node.Operator, left, right, env)

	case *ast.IfExpression:
		return evalIfExpression(node, env)

	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}

	case *ast.Boolean:
		return boolean(node.Value)

	}

	return Null
}

func evalProgram(program *ast.Program, env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range program.Statements {
		result = Eval(statement, env)

		switch result := result.(type) {
		case *object.ReturnValue:
			return result.Value
		case *object.Error:
			return result
		}
	}

	return result
}

func evalBlockStatement(bs *ast.BlockStatement, env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range bs.Statements {
		result = Eval(statement, env)

		if result != nil {

			typ := result.Type()
			if typ == object.RETURN_VALUE || typ == object.ERROR {
				return result
			}
		}
	}

	return result
}

func evalLetStatement(ls *ast.LetStatement, env *object.Environment) object.Object {
	val := Eval(ls.Value, env)
	if isError(val) {
		return val
	}

	env.Set(ls.Name.Value, val)
	return val
}

func evalIdentifier(ident *ast.Identifier, env *object.Environment) object.Object {
	val, ok := env.Get(ident.Value)
	if !ok {
		return newError("identifier not found: %s", ident.Value)
	}

	return val
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

func evalPrefixExpression(op string, obj object.Object, env *object.Environment) object.Object {

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

func isError(obj object.Object) bool {
	return obj.Type() == object.ERROR
}

func evalExpressions(exps []ast.Expression, env *object.Environment) []object.Object {
	var result []object.Object

	for _, e := range exps {
		evaluated := Eval(e, env)
		if isError(evaluated) {
			return []object.Object{evaluated}
		}
		result = append(result, evaluated)
	}

	return result
}

func applyFunction(fn object.Object, args []object.Object) object.Object {
	function, ok := fn.(*object.Function)
	if !ok { 
		return newError("not a function: %s", fn.Type())
	}

	envWithArgs := function.Env.NewChild()

	for paramIdx, param := range function.Parameters {
		envWithArgs.Set(param.Value, args[paramIdx])
	}

	evaluated := Eval(function.Body, envWithArgs)
	return unwrapReturnValue(evaluated)
}

func unwrapReturnValue(val object.Object) object.Object {
	if returnValue, ok := val.(*object.ReturnValue); ok {
		return returnValue.Value
	}

	return val
}