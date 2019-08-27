package eval

import (
	"fmt"
	"monkey-go/object"
	"monkey-go/token"
)

func evalInfixExpression(op string, l, r object.Object, env *object.Environment) object.Object {
	switch {
	case l.Type() != r.Type():
		return newError("type mismatch: %s %s %s", l.Type(), op, r.Type())

	case l.Type() == object.BOOLEAN:
		return evalBoolInfix(op, l, r, env)

	case l.Type() == object.INTEGER:
		return evalIntegerInfix(op, l, r, env)

	case l.Type() == object.STRING:
		return evalStringInfix(op, l, r, env)

	default:
		return newError("unknown operator: %s %s %s", l.Type(), op, r.Type())
	}
}

func evalBoolInfix(op string, l, r object.Object, env *object.Environment) object.Object {
	if op == token.EQ {
		return boolean(l == r)
	}

	if op == token.NE {
		return boolean(l != r)
	}

	return newError("unknown operator: %s %s %s", l.Type(), op, r.Type())
}

func evalIntegerInfix(op string, l, r object.Object, env *object.Environment) object.Object {

	leftVal := l.(*object.Integer).Value
	rightVal := r.(*object.Integer).Value

	if evaluator, ok := integerInfixes[op]; ok {
		return doEval(evaluator, leftVal, rightVal)
	}

	return Null
}

func evalStringInfix(op string, l, r object.Object, env *object.Environment) object.Object {

	leftVal := l.(*object.String).Value
	rightVal := r.(*object.String).Value

	if evaluator, ok := stringInfixes[op]; ok {
		return doStringEval(evaluator, leftVal, rightVal)
	}

	return Null
}

func doStringEval(e stringInfixEvaluator, l, r string) object.Object {
	result := e(l, r)

	switch v := result.(type) {
	case bool:
		return boolean(v)

	case string:
		return &object.String{Value: v}

	default:
		return Null
	}
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

type stringInfixEvaluator func(l, r string) interface{}

var stringInfixes = map[string]stringInfixEvaluator{
	"+":  concat,
	"==": strcmp,
	"!=": notstrcmp,
}

func concat(l, r string) interface{}    { return fmt.Sprintf("%s%s", l, r) }
func strcmp(l, r string) interface{}    { return l == r }
func notstrcmp(l, r string) interface{} { return l != r }
