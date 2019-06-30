package object

import "fmt"

type Integer struct {
	Value int
}

func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }

func (i *Integer) Type() Type { return INTEGER }

