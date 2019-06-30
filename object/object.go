package object

type Type int

type Object interface {
	Type() Type
	Inspect() string
	Truthy() bool
}

const (
	INTEGER = iota + 1
	BOOLEAN
	NULL
)
