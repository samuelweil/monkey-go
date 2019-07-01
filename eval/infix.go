package eval

import (
	"monkey-go/object"
	"monkey-go/token"
)

func evalInfixExpression(op string, l, r object.Object) object.Object {
	switch {
	case l.Type() != r.Type():
		return newError("type mismatch: %s %s %s", l.Type(), op, r.Type())

	case l.Type() == object.BOOLEAN:
		return evalBoolInfix(op, l, r)

	case l.Type() == object.INTEGER:
		return evalIntegerInfix(op, l, r)

	default:
		return newError("unknown operator: %s %s %s", l.Type(), op, r.Type())
	}
}

func evalBoolInfix(op string, l, r object.Object) object.Object {
	if op == token.EQ {
		return boolean(l == r)
	}

	if op == token.NE {
		return boolean(l != r)
	}

	return newError("unknown operator: %s %s %s", l.Type(), op, r.Type())
}

func evalIntegerInfix(op string, l, r object.Object) object.Object {

	leftVal := l.(*object.Integer).Value
	rightVal := r.(*object.Integer).Value

	if evaluator, ok := integerInfixes[op]; ok {
		return doEval(evaluator, leftVal, rightVal)
	}

	return Null
}

func doEval(e intInfixEvaluator, l, r int) object.Object {

	result := e(l, r)

	switch v := result.(type) {

	case bool:
		return boolean(v)

	case int:
		return &object.Integer{Value: v}

	default:
		return Null
	}
}

type intInfixEvaluator func(l, r int) interface{}

var integerInfixes = map[string]intInfixEvaluator{
	"+":  add,
	"-":  subtract,
	"*":  multiply,
	"/":  divide,
	"==": equal,
	"!=": notEqual,
	"<":  less,
	">":  greater,
}

func add(l, r int) interface{}      { return l + r }
func subtract(l, r int) interface{} { return l - r }
func multiply(l, r int) interface{} { return l * r }
func divide(l, r int) interface{}   { return l / r }

func equal(l, r int) interface{}    { return l == r }
func notEqual(l, r int) interface{} { return l != r }
func less(l, r int) interface{}     { return l < r }
func greater(l, r int) interface{}  { return l > r }
