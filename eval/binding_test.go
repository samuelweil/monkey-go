package eval

import (
	"testing"

	"github.com/samuelweil/go-tools/testing/check"
)

func TestLetStatements(t *testing.T) {
	check := check.New(t)
	
	tests := []struct {
		input string
		exp   int
	}{
		{"let a = 10; a", 10},
		{"let a = 5 * 5; a", 25},
		{"let a = 5; let b = a; b", 5},
		{"let a = 5; let b = a; let c = a + b + 5; c", 15},
	}

	for _, tt := range tests {
		result := testEval(tt.input)
		check.NoError(testIntegerObject(result, tt.exp))
	}
}
