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

var integerInfixes = map[string]intInfixEvaluator {
	"+": add,
	"-": subtract,
	"*": multiply,
	"/": divide,
	"==": equal,
	"!=": notEqual,
	"<": less,
	">": greater,
}

func add(l, r int) interface{} { return l + r }
func subtract(l, r int) interface{} { return l - r}
func multiply(l, r int) interface{} {return l * r }
func divide(l, r int) interface{} { return l / r}

func equal(l, r int) interface{} { return l == r }
func notEqual(l, r int) interface{} { return l != r}
func less(l, r int) interface{} { return l < r}
func greater(l, r int) interface{} { return l > r}