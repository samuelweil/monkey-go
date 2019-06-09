package token

const (
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
	BANG     = "!"
)

func Assign() Token {
	return FromStr(ASSIGN)
}
func Plus() Token {
	return FromStr(PLUS)
}

func Minus() Token {
	return FromStr(MINUS)
}

func Asterisk() Token {
	return FromStr(ASTERISK)
}

func Slash() Token {
	return FromStr(SLASH)
}

func Bang() Token {
	return FromStr(BANG)
}

var operators = map[string]Token{
	ASSIGN:   Assign(),
	PLUS:     Plus(),
	MINUS:    Minus(),
	ASTERISK: Asterisk(),
	SLASH:    Slash(),
	BANG:     Bang(),
}

func Operator(c byte) Token {
	return operators[string(c)]
}

func IsOperator(c byte) bool {
	_, b := operators[string(c)]
	return b
}
