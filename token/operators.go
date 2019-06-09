package token

const (
	// Operators
	ASSIGN = "="
	PLUS   = "+"
)

func Plus() Token {
	return FromStr(PLUS)
}

func Assign() Token {
	return FromStr(ASSIGN)
}
