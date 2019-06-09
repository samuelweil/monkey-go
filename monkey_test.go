package main

import "testing"

func assertEq(t *testing.T, expected, got interface{}) {
	if got != expected {
		t.Errorf("Expected %v, got %v\n", expected, got)
	}
}

func TestHello(t *testing.T) {
	assertEq(t, "Hello World", hello())
}
