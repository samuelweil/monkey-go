package token

const (
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
	BANG     = "!"
	AMP      = "&"
	PIPE     = "|"

	EQ = "=="
	NE = "!="
	LE = "<="
	GE = ">="
	LT = "<"
	GT = ">"

	AND = "&&"
	OR  = "||"
)

var Operators = Set{
	ASSIGN:   Assign(),
	PLUS:     Plus(),
	MINUS:    Minus(),
	ASTERISK: Asterisk(),
	SLASH:    Slash(),
	BANG:     Bang(),
	EQ:       Eq(),
	NE:       NotEq(),
	LE:       LessEq(),
	GE:       GreatEq(),
	LT:       LessThan(),
	GT:       GreaterThan(),
	AND:      And(),
	OR:       Or(),
	AMP:      Ampersand(),
	PIPE:     Pipe(),
}

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

func Eq() Token {
	return FromStr(EQ)
}

func NotEq() Token {
	return FromStr(NE)
}

func LessEq() Token {
	return FromStr(LE)
}

func GreatEq() Token {
	return FromStr(GE)
}

func LessThan() Token {
	return FromStr(LT)
}

func GreaterThan() Token {
	return FromStr(GT)
}

func And() Token {
	return FromStr(AND)
}

func Or() Token {
	return FromStr(OR)
}

func Ampersand() Token {
	return FromStr(AMP)
}

func Pipe() Token {
	return FromStr(PIPE)
}
