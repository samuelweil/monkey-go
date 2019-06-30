package object

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() Type { return RETURN_VALUE }
func (rv *ReturnValue) Inspect() string { return rv.Value.Inspect() }
func (rv *ReturnValue) Truthy() bool { return rv.Value.Truthy() }