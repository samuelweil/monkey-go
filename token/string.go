package token

const (
	STRING = "STRING"
)

func String(s string) Token {
	return Token{
		Type:    STRING,
		Literal: s,
	}
}
