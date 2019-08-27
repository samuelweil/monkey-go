package object

import "fmt"

type String struct {
	Value string
}

func (s *String) Inspect() string { return fmt.Sprintf("\"%s\"", s.Value) }
func (s *String) Type() Type      { return STRING }
func (s *String) Truthy() bool    { return s.Value != "" }
