package eval

import "monkey-go/object"

var prefixes = map[string]func(object.Object) object.Object{
	"!": func(obj object.Object) object.Object {
		switch obj {
		case False, Null:
			return True
		default:
			return False
		}
	},
}
