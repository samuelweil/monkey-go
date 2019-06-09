package token

const (
	FUNCTION = "FUNC"
	LET      = "LET"
)

var KeyWords = map[string]Token{
	"fn":  Function(),
	"let": Let(),
}

func Function() Token {
	return FromStr(FUNCTION)
}
