package testing

type Formatter func(fmt string, args ...interface{})

type Tester struct {
	formatter Formatter
}

func (t *Tester) Eq(got, expected interface{}) bool {
	ok := got == expected
	if !ok {
		t.formatter("Expected %v, Got %v", expected, got)
	}

	return ok
}

func (t *Tester) NotEq(got, expected interface{}) bool {
	ok := (got != expected)
	if !ok {
		t.formatter("Two items shouldn't be equal")
	}

	return ok
}

func New(f Formatter) Tester {
	return Tester{formatter: f}
}
