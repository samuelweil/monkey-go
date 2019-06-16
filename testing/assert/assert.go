package assert

import (
	mnktest "monkey-go/testing"
	"testing"
)

func New(t *testing.T) mnktest.Tester {
	return mnktest.New(t.Fatalf)
}
