package lexer

import "monkey-go/token"

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

type RuleTokenizer struct {
	rule        func(b byte) bool
	constructor func(s string) token.Token
}

func (r RuleTokenizer) Check(b byte) bool {
	return r.rule(b)
}

func (r RuleTokenizer) GetToken(s string) (token.Token, string) {

	for i := 0; i < len(s); i++ {
		if !r.Check(s[i]) {
			return r.constructor(s[:i]), s[i:]
		}
	}

	return r.constructor(s), ""

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
