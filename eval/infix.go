package eval

import "monkey-go/object"

func evalInfixExpression(op string, l, r object.Object) object.Object {

	if _, ok := l.(*object.Integer); !ok {
		return Null
	}

	if _, ok := r.(*object.Integer); !ok {
		return Null
	}

	return evalIntegerInfix(op, l, r)
}

func evalIntegerInfix(op string, l, r object.Object) object.Object {

	leftVal := l.(*object.Integer).Value
	rightVal := r.(*object.Integer).Value

	switch op {
	case "+":
		return &object.Integer{Value: leftVal + rightVal}

	case "-":
		return &object.Integer{Value: leftVal - rightVal}

	case "*":
		return &object.Integer{Value: leftVal * rightVal}

	case "/":
		return &object.Integer{Value: leftVal / rightVal}

	default:
		return Null
	}
}