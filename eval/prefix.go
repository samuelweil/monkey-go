package eval

import "monkey-go/object"

var prefixEvals = map[string]func(object.Object) object.Object{
	"!": bang,
	"-": minus,
}

func bang(obj object.Object) object.Object {
	switch obj {
	case False, Null:
		return True
	default:
		return False
	}
}

func minus(obj object.Object) object.Object {
	val, ok := obj.(*object.Integer)
	if !ok {
		return Null
	}

	return &object.Integer{Value: -val.Value}
}
