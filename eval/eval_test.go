package eval

import (
	"fmt"
	"testing"

	"monkey-go/object"
	"monkey-go/parser"
	
	"github.com/samuelweil/go-tools/testing/check"
)

func testEval(inp string) object.Object {
	p := parser.New(inp)
	program := p.ParseProgram()

	return Eval(program)
}

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

func TestBooleanEval(t *testing.T) {

	check := check.New(t)

	tests := []struct{
		input string
		expected bool
	} {
		{"true", true},
		{"false", false},
	}

	for _, tt := range tests {
		result := testEval(tt.input)
		check.Nil(testBooleanObject(result, tt.expected))
	}
}

func testBooleanObject(obj object.Object, exp bool) error {
	result, ok := obj.(*object.Boolean)
	if !ok {
		return fmt.Errorf("object is not a Boolean. Got %T (%+v)", obj, obj)
	}

	if bool(*result) != exp {
		return fmt.Errorf("Boolean is not %t. Got %t", bool(*result), exp)
	}

	return nil
}