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
		input    string
		expected int
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2", 16},
		{"-50 + 100 - 50", 0},
		{"5 * 2 + 10", 20},
		{"20 + 2 * -10", 0},
		{"50 / 2 * 2 + 10", 60},
		{"2 * (5 + 10)", 30},
		{"3 * 3 * 3 + 10", 37},
		{"3 * (3 * 3) + 10", 37},
		{"3 * 3 * (3 + 10)", 117},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
	}

	for _, tt := range tests {
		result := testEval(tt.input)
		check.NoError(testIntegerObject(result, tt.expected))
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

	tests := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 < 1", false},
		{"1 > 1", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"1 != 2", true},
		{"true == true", true},
		{"true != true", false},
		{"false == false", true},
		{"false != false", false},
		{"true == false", false},
		{"true != false", true},
		{"(1 < 2) == true", true},
		{"(1 < 2) == false", false},
		{"(1 > 2) == true", false},
		{"(1 > 2) == false", true},
	}

	for _, tt := range tests {
		result := testEval(tt.input)
		check.NoError(testBooleanObject(result, tt.expected))
	}
}

func testBooleanObject(obj object.Object, exp bool) error {
	result, ok := obj.(*object.Boolean)
	if !ok {
		return fmt.Errorf("object is not a Boolean. Got %T (%+v)", obj, obj)
	}

	if result.Value != exp {
		return fmt.Errorf("Boolean is not %t. Got %t", result.Value, exp)
	}

	return nil
}

func TestBangOperator(t *testing.T) {

	check := check.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		{"!true", false},
		{"!false", true},
		{"!5", false},
		{"!!true", true},
		{"!!false", false},
		{"!!5", true},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		check.NoError(testBooleanObject(evaluated, tt.expected))
	}
}
