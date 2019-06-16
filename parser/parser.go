package parser

import (
	"fmt"
	"monkey-go/ast"
	"monkey-go/lexer"
	"monkey-go/token"
	"strconv"
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

type Parser struct {
	l *lexer.Lexer

	currentToken token.Token
	peekToken    token.Token

	errors []string

	prefixParsers map[token.Type]prefixParseFn
	infixParsers  map[token.Type]infixParseFn
}

func New(s string) *Parser {
	p := &Parser{
		l:      lexer.New(s),
		errors: []string{},
	}

	p.prefixParsers = make(map[token.Type]prefixParseFn)
	p.infixParsers = make(map[token.Type]infixParseFn)

	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
}

func (p *Parser) registerPrefix(typ token.Type, fn prefixParseFn) {
	p.prefixParsers[typ] = fn
}

func (p *Parser) registerInfix(typ token.Type, fn infixParseFn) {
	p.infixParsers[typ] = fn
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {

	program := new(ast.Program)

	for p.currentToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currentToken.Type {
	case token.LET:
		return p.parseLetStatement()

	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.currentToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	for !p.currentTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) expectPeek(typ token.Type) bool {
	if p.peekTokenIs(typ) {
		p.nextToken()
		return true
	}

	p.peekError(typ)
	return false
}

func (p *Parser) peekTokenIs(typ token.Type) bool {
	return p.peekToken.Type == typ
}

func (p *Parser) currentTokenIs(typ token.Type) bool {
	return p.currentToken.Type == typ
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.Type) {
	msg := fmt.Sprintf("expected next token to be %s, got %s", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.currentToken}

	p.nextToken()

	for !p.currentTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.currentToken}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParsers[p.currentToken.Type]
	if prefix == nil {
		return nil
	}

	return prefix()
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.currentToken}

	value, err := strconv.Atoi(p.currentToken.Literal)
	if err != nil {
		msg := fmt.Sprintf("Could not parse %q as integer", p.currentToken.Literal)
		p.errors = append(p.errors, msg)
	}

	lit.Value = value
	return lit

}
