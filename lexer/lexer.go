package lexer

import "monkey-go/token"

type Lexer struct {
	tokenizers []tokenizer
	input      string
}

func New(input string) *Lexer {
	return &Lexer{
		input: input,
		tokenizers: []tokenizer{
			SetTokenizer{token.Operators},
			SetTokenizer{token.Delimiters},
			&WhileTokenizer{
				while:       isDigit,
				constructor: token.Int,
			},
			IdentTokenizer{},
		},
	}
}

func (l *Lexer) NextToken() token.Token {
	l.input = skipWhiteSpace(l.input)

	if len(l.input) == 0 {
		return token.Eof()
	}

	for _, t := range l.tokenizers {
		if t.Check(l.input[0]) {
			tok, remain := t.GetToken(l.input)
			l.input = remain
			return tok
		}
	}

	return token.Illegal()
}

func skipWhiteSpace(s string) string {
	for i := 0; i < len(s); i++ {
		if !isWhiteSpace(s[i]) {
			return s[i:]
		}
	}

	return ""
}
