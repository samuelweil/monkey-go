package object

type Environment struct {
	store map[string]Object
	parent *Environment
}

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{
		store: s,
		parent: nil,
	}
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]

	if !ok && e.parent != nil {
		return e.parent.Get(name)
	}

	return obj, ok

}

func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

func (e  *Environment) NewChild() *Environment {
	return &Environment{
		store: make(map[string]Object),
		parent: e,
	}
}