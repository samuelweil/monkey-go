package lexer

import "monkey-go/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(s string) Lexer {
	l := Lexer{input: s}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func fromChar(c byte) token.Token {
	return token.FromStr(string(c))
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=', '+', ',', ';', '(', ')', '{', '}', '[', ']':
		tok = fromChar(l.ch)
	case 0:
		tok = token.Token{Type: token.EOF}
	}

	l.readChar()
	return tok
}
