package token

type Type string

type Token struct {
	Type    Type
	Literal string
}

func FromStr(s string) Token {
	return Token{Type(s), s}
}

func Let() Token {
	return Token{
		Type: LET,
	}
}

func Ident(id string) Token {
	return Token{
		Type:    IDENT,
		Literal: id,
	}
}

func Assign() Token {
	return FromStr(ASSIGN)
}

func Int(s string) Token {
	return Token{
		Type:    INT,
		Literal: s,
	}
}

func SemiColon() Token {
	return FromStr(SEMICOLON)
}

func Function() Token {
	return FromStr(FUNCTION)
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

func Comma() Token {
	return FromStr(COMMA)
}

func Plus() Token {
	return FromStr(PLUS)
}

func Eof() Token {
	return Token{EOF, ""}
}

func Illegal() Token {
	return Token{
		Type: ILLEGAL,
	}
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + Literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	LBRACK = "["
	RBRACK = "]"

	// Keywords
	FUNCTION = "FUNC"
	LET      = "LET"
)

var KeyWords = map[string]Token{
	"fn":  Function(),
	"let": Let(),
}
