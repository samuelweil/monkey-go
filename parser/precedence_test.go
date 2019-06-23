package parser

import (
	"testing"

	"github.com/samuelweil/go-tools/testing/check"
)

type precedenceTests []struct {
	input    string
	expected string
}

func TestOperatorPrecedenceParsing(t *testing.T) {
	tests := precedenceTests{
		{
			"-a * b",
			"((-a) * b)",
		},
		{
			"!-a",
			"(!(-a))",
		},
		{
			"a + b + c",
			"((a + b) + c)",
		},
		{
			"a * b + c",
			"((a * b) + c)",
		},
		{
			"3 + 4; -5 * 5",
			"(3 + 4)((-5) * 5)",
		},
		{
			"3 + 4 * 5 == 3 * 1 + 4 * 5",
			"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
		},
	}

	testPrecedence(t, tests)
}

func TestBooleanPrecedence(t *testing.T) {
	tests := precedenceTests{
		{
			"true",
			"true",
		},
		{
			"false",
			"false",
		},
		{
			"3 > 5 == false",
			"((3 > 5) == false)",
		},
		{
			"3 < 5 == true",
			"((3 < 5) == true)",
		},
	}

	testPrecedence(t, tests)
}

func testPrecedence(t *testing.T, tests precedenceTests) {

	check := check.New(t)

	for _, tt := range tests {
		p := New(tt.input)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		check.Eq(program.String(), tt.expected)
	}
}
