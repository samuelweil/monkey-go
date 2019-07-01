package eval

import (
	"monkey-go/object"
	"testing"

	"github.com/samuelweil/go-tools/testing/check"
)

func TestErrorHandling(t *testing.T) {
	check := check.New(t)

	tests := []struct {
		input  string
		expMsg string
	}{
		{
			"5 + true",
			"type mismatch: INT + BOOL",
		},
		{
			"5 + true; 5;",
			"type mismatch: INT + BOOL",
		},
		{
			"-true",
			"unknown operator: -BOOL",
		},
		{
			"true + false",
			"unknown operator: BOOL + BOOL",
		},
		{
			"5; true + false; 17",
			"unknown operator: BOOL + BOOL",
		},
		{
			"if (10 > 1) { true + false; }",
			"unknown operator: BOOL + BOOL",
		},
		{
			`if (10 > 1) {
				if (true) {
					ret true + false
				}
			}`,
			"unknown operator: BOOL + BOOL",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		err, ok := evaluated.(*object.Error)
		if check.True(ok, "Expected an error object. Got %T (%+v)", evaluated, evaluated) {
			check.Eq(err.Message, tt.expMsg)
		}
	}
}
