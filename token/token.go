package token

import "strings"

type Type string

type Token struct {
	Type    Type
	Literal string
}

func FromStr(s string) Token {
	return Token{Type(s), strings.ToLower(s)}
}

func Ident(id string) Token {
	return Token{
		Type:    IDENT,
		Literal: id,
	}
}

func Int(s string) Token {
	return Token{
		Type:    INT,
		Literal: s,
	}
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

	// Values
	IDENT = "IDENT"
	INT   = "INT"
)
