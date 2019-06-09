package token

const (
	FUNCTION = "FUNC"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	FOR      = "FOR"
)

var KeyWords = map[string]Token{
	"fn":    Function(),
	"let":   Let(),
	"true":  True(),
	"false": False(),
	"if":    If(),
	"else":  Else(),
	"ret":   Return(),
	"for":   For(),
}

func Function() Token {
	return FromStr(FUNCTION)
}

func Let() Token {
	return FromStr(LET)
}

func True() Token {
	return FromStr(TRUE)
}

func False() Token {
	return FromStr(FALSE)
}

func If() Token {
	return FromStr(IF)
}

func Else() Token {
	return FromStr(ELSE)
}

func Return() Token {
	return FromStr(RETURN)
}

func For() Token {
	return FromStr(FOR)
}
