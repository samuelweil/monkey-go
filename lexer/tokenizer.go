package lexer

import (
	"monkey-go/token"
)

type tokenizer interface {
	Check(b byte) bool
	GetToken(s string) (token.Token, string)
}

type SetTokenizer struct {
	tokenSet token.Set
}

func (t SetTokenizer) Check(b byte) bool {
	return t.checkString(string(b))
}

func (t SetTokenizer) GetToken(input string) (token.Token, string) {
	i := 1
	n := len(input)
	for ; t.checkString(input[:i]); i++ {
		if i == n {
			return t.tokenSet[input], ""
		}
	}

	i--

	tk := t.tokenSet[input[:i]]
	return tk, input[i:]
}

func (t SetTokenizer) checkString(s string) bool {
	_, ok := t.tokenSet[s]
	return ok
}

type WhileTokenizer struct {
	while       func(byte) bool
	constructor func(string) token.Token
}

func (w *WhileTokenizer) Check(b byte) bool {
	return w.while(b)
}

func (w *WhileTokenizer) GetToken(s string) (token.Token, string) {

	var i int
	for i = 0; i < len(s) && w.while(s[i]); i++ {
	}

	return w.constructor(s[:i]), s[i:]

}

type BetweenTokenizer struct {
	trigger     byte
	constructor func(string) token.Token
}

func (bt *BetweenTokenizer) Check(b byte) bool {
	return b == bt.trigger
}

func (bt *BetweenTokenizer) GetToken(s string) (token.Token, string) {
	var i int

	for i = 1; i < len(s) && s[i] != bt.trigger; i++ {
	}

	return bt.constructor(s[1:i]), s[i+1:]
}

type UntilTokenizer struct {
	from        func(b byte) bool
	until       func(b byte) bool
	constructor func(s string) token.Token
}

func (u *UntilTokenizer) Check(b byte) bool {
	return u.from(b)
}

func (u *UntilTokenizer) GetToken(s string) (token.Token, string) {
	var i int

	for i = 0; i < len(s) && !u.until(s[i]); i++ {
	}

	return u.constructor(s[:i]), s[i:]
}

type IdentTokenizer struct{}

func (t IdentTokenizer) Check(b byte) bool {
	return isLetter(b) || (b == '_')
}

func (t IdentTokenizer) GetToken(s string) (token.Token, string) {

	name, remain := s, ""

	for i := 0; i < len(s); i++ {
		if !t.Check(s[i]) {
			name, remain = s[:i], s[i:]
			break
		}
	}

	if kw, ok := token.KeyWords[name]; ok {
		return kw, remain
	}

	return token.Ident(name), remain
}
