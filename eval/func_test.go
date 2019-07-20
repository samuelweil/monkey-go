package eval

import (
	"monkey-go/object"
	"testing"

	"github.com/samuelweil/go-tools/testing/assert"
	"github.com/samuelweil/go-tools/testing/check"
)

func TestFunctionObject(t *testing.T) {

	assert := assert.New(t)

	input := "fn(x) { x + 2; };"

	evaluated := testEval(input)
	fn, ok := evaluated.(*object.Function)
	assert.True(ok, "object is not Function. Got %T (%+v)", evaluated, evaluated)

	assert.Eq(len(fn.Parameters), 1)
	assert.Eq(fn.Parameters[0].String(), "x")

	expectedBody := "(x + 2)"

	assert.Eq(fn.Body.String(), expectedBody)
}

func TestFunctionApplication(t *testing.T) {

	check := check.New(t)

	tests := []struct {
		input    string
		expected int
	}{
		{"let identity = fn(x) { x; }; identity(5);", 5},
		{"let identity = fn(x) { ret x; }; identity(5);", 5},
		{"let double = fn(x) { x * 2; }; double(5);", 10},
		{"let add = fn(x, y) { x + y; }; add(5, 5);", 10},
		{"let add = fn(x, y) { x + y; }; add(5 + 5, add(5, 5));", 20},
		{"fn(x) { x; }(5)", 5},
	}

	for _, tt := range tests {
		check.NoError(testIntegerObject(testEval(tt.input), tt.expected))
	}
}
