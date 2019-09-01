package eval

import (
	"monkey-go/object"
)

var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{
		Fn: length,
	},
	"exit": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			return nil
		},
	},
}

func length(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("len takes 1 arguments, got %d", len(args))
	}

	switch arg := args[0].(type) {
	case *object.String:
		return &object.Integer{Value: len(arg.Value)}

	case *object.Array:
		return &object.Integer{Value: len(arg.Elements)}

	default:
		return newError("len does not support arguments of type %s", args[0].Type())
	}
}
