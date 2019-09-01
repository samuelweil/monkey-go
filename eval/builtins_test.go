package eval

import (
	"monkey-go/object"
	"testing"

	"github.com/samuelweil/go-tools/testing/assert"
)

func TestBuiltinFuncs(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`len("")`, 0},
		{`len("four")`, 4},
		{`len("hello, world")`, 12},
		{`len(1)`, "len does not support arguments of type INT"},
		{`len("one", "two")`, "len takes 1 arguments, got 2"},
	}

	assert := assert.New(t)

	for _, tt := range tests {

		evaluated := testEval(tt.input)

		switch expected := tt.expected.(type) {
		case int:
			assert.NoError(testIntegerObject(evaluated, expected))

		case string:
			errObj, ok := evaluated.(*object.Error)
			assert.True(ok, "object is not Error. got %T, (%+v)", evaluated, evaluated)
			assert.Eq(errObj.Message, expected)
		}
	}
}
