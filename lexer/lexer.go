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
	l.readPosition++
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}

	return l.input[l.readPosition]
}

func (l *Lexer) NextToken() (tok token.Token) {
	l.skipWhiteSpace()

	switch {
	case l.ch == 0:
		tok = token.Eof()

	case isDigit(l.ch):
		tok = l.readNumber()
		return

	case isValidIdentChar(l.ch):
		literal := l.readIdentifier()

		if kw, ok := token.KeyWords[literal]; ok {
			tok = kw
		} else {
			tok = token.Ident(literal)
		}

		return

	case token.IsOperator(l.ch):
		tok = token.Operator(l.ch)

		// Check for a 2-char operator
		op := string(l.ch) + string(l.peekChar())
		if token.IsMultiCharOperator(op) {
			tok = token.MultiCharOperator(op)
			l.readChar()
		}

	case token.IsDelimiter(l.ch):
		tok = token.Delimiter(l.ch)

	default:
		tok = token.Illegal()
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
