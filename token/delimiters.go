package token

const (
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	LBRACK = "["
	RBRACK = "]"
)

func SemiColon() Token {
	return FromStr(SEMICOLON)
}

func Comma() Token {
	return FromStr(COMMA)
}

func LParen() Token {
	return FromStr(LPAREN)
}

func RParen() Token {
	return FromStr(RPAREN)
}

func LBrace() Token {
	return FromStr(LBRACE)
}

func RBrace() Token {
	return FromStr(RBRACE)
}

func LBrack() Token {
	return FromStr(LBRACK)
}

func RBrack() Token {
	return FromStr(RBRACK)
}
