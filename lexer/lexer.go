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

	l.skipWhiteSpace()

	switch l.ch {
	case '=', '+', ',', ';', '(', ')', '{', '}', '[', ']':
		tok = fromChar(l.ch)
	case 0:
		tok = token.Token{Type: token.EOF}
	default:
		switch {

		case isValidIdentChar(l.ch):
			literal := l.readIdentifier()

			if keywordToken, ok := token.KeyWords[literal]; ok {
				tok = keywordToken
			} else {
				tok = token.Ident(literal)
			}
			return tok

		case isDigit(l.ch):
			return l.readNumber()

		default:
			tok = token.Illegal()
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier() string {
	start := l.position

	for isValidIdentChar(l.ch) {
		l.readChar()
	}

	return l.input[start:l.position]
}

func (l *Lexer) readNumber() token.Token {

	start := l.position

	for isDigit(l.ch) {
		l.readChar()
	}

	return token.Int(l.input[start:l.position])
}

func (l *Lexer) skipWhiteSpace() {
	for isWhiteSpace(l.ch) {
		l.readChar()
	}
}

func isValidIdentChar(c byte) bool {
	return isLetter(c) || (c == '_')
}
