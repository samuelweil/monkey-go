package assert

import "testing"

type Asserter struct {
	t *testing.T
}

func New(t *testing.T) *Asserter {
	return &Asserter{t: t}
}

func (a *Asserter) Eq(expected, got interface{}) bool {
	ok := expected == got
	if !ok {
		a.t.Errorf("Expected %v, Got %v", expected, got)
	}
	return ok
}
