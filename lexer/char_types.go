package lexer

func isLetter(c byte) bool {
	return isLowerCase(c) || isUpperCase(c)
}

func isLowerCase(c byte) bool {
	return ('a' <= c && c <= 'z')
}

func isUpperCase(c byte) bool {
	return ('A' <= c && c <= 'Z')
}

func isWhiteSpace(c byte) bool {
	switch c {
	case ' ', '\t', '\n', '\r':
		return true
	default:
		return false
	}
}

func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}
