package eval

import (
	"fmt"
	"testing"

	"monkey-go/object"
	"monkey-go/parser"
	
	"github.com/samuelweil/go-tools/testing/check"
)

func TestIntegerEval(t *testing.T) {

	check := check.New(t)

	tests := []struct {
		input string
		expected int
	} {
		{"5", 5},
		{"10", 10},
	}

	for _, tt := range tests {
		result := testEval(tt.input)
		check.Nil(testIntegerObject(result, tt.expected))	
	}
}

func testEval(inp string) object.Object {
	p := parser.New(inp)
	program := p.ParseProgram()

	return Eval(program)
}

func testIntegerObject(obj object.Object, expected int) error {
	result, ok := obj.(*object.Integer)
	if !ok {
		return fmt.Errorf("object is not integer. Got %T (%+v)", obj, obj)
	}

	if result.Value != expected {
		return fmt.Errorf("Integer value %d expected. Got %d", expected, result.Value)
	}

	return nil
}