package token

const (
	ASSIGN = "="
	PLUS   = "+"
)

func Plus() Token {
	return FromStr(PLUS)
}

func Assign() Token {
	return FromStr(ASSIGN)
}

var operators = map[string]Token{
	PLUS:   Plus(),
	ASSIGN: Assign(),
}

func Operator(c byte) Token {
	return operators[string(c)]
}

func IsOperator(c byte) bool {
	_, b := operators[string(c)]
	return b
}
