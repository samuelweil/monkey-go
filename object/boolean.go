package object

import (
	"fmt"
)

type Boolean struct {
	Value bool
}

func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }
func (b *Boolean) Type() Type { return BOOLEAN }
func (b *Boolean) Truthy() bool { return b.Value }
