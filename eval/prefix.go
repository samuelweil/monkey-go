package eval

import "monkey-go/object"

var prefixEvals = map[string]func(object.Object) object.Object{
	"!": bang,
	"-": minus,
}

func bang(obj object.Object) object.Object {
	return boolean(!obj.Truthy())
}

func minus(obj object.Object) object.Object {
	val, ok := obj.(*object.Integer)
	if !ok {
		return newError("unknown operator: -%s", obj.Type())
	}

	return &object.Integer{Value: -val.Value}
}
