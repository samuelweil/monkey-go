package object

type Null struct{}

func (n *Null) Inspect() string { return "null" }
func (n *Null) Type() Type { return NULL }
func (n *Null) Truthy() bool { return false }