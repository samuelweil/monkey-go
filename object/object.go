package object

type Type string

type Object interface {
	Type() Type
	Inspect() string
	Truthy() bool
}

const (
	INTEGER      = "INT"
	BOOLEAN      = "BOOL"
	NULL         = "NULL"
	RETURN_VALUE = "RETURN"
	ERROR        = "ERROR"
	FUNCTION     = "FUNCTION"
)
