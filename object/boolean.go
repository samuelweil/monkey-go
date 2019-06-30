package object

import ("fmt")

type Boolean bool

func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", bool(*b)) }

func (b *Boolean) Type() Type { return BOOLEAN }

