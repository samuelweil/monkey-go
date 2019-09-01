package eval

import (
	"testing"

	"monkey-go/object"

	"github.com/samuelweil/go-tools/testing/assert"
	"github.com/samuelweil/go-tools/testing/check"
)

func TestArrayEval(t *testing.T) {

	assert := assert.New(t)

	input := "[1, 2, 3]"

	evaluated := testEval(input)
	result, ok := evaluated.(*object.Array)

	assert.True(ok, "object is not Array. Got %T (%+v)", evaluated, evaluated)
	assert.Eq(len(result.Elements), 3)

	assert.NoError(testIntegerObject(result.Elements[0], 1))
	assert.NoError(testIntegerObject(result.Elements[1], 2))
	assert.NoError(testIntegerObject(result.Elements[2], 3))
}

func TestArrayIndex(t *testing.T) {
	check := check.New(t)

	tests := []struct {
		input string
		exp   interface{}
	}{
		{
			"[1, 2, 3][0]",
			1,
		},
		{
			"['hello', 'world'][1]",
			"world",
		},
		{
			"let i = 0; [1][i];",
			1,
		},
		{
			"[1 , 2, 3][1 + 1]",
			3,
		}, {
			"let arr = [1, 2, 3]; arr[1]",
			2,
		},
		{
			"let arr = [1,2,3]; arr[0] + arr[1] + arr[2]",
			6,
		},
		{
			"let arr = [fn(x, y){ x + y}]; arr[arr[0](1, -1)](1, 2)",
			3,
		},
		{
			"[1, 2, 3][-1]",
			"index out of bounds",
		},
		{
			"[1, 2, 3][4]",
			"index out of bounds",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		switch e := tt.exp.(type) {
		case int:
			check.NoError(testIntegerObject(evaluated, e))

		case string:
			if err, ok := evaluated.(*object.Error); ok {
				check.Eq(err.Message, e)
			} else {
				check.NoError(testStringObject(evaluated, e))
			}
		}
	}
}
