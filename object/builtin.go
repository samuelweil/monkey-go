package object

type BuiltinFunc func(args ...Object) Object
type Builtin struct {
	Fn BuiltinFunc
}

func (b *Builtin) Type() Type      { return BUILTIN }
func (b *Builtin) Inspect() string { return "builtin function" }
func (b *Builtin) Truthy() bool    { return true }
