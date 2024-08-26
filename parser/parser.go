package parser

import (
	"fmt"
	"monky/ast"
	"monky/lexer"
	"monky/token"
)

const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER //> or <
	SUM         //+
	PRODUCT     //*
	PREFIX      //!x or -x
	CAll        //myFunction(x)
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(expression ast.Expression) ast.Expression
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	errors    []string

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParserProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	for p.curToken.Type != token.EOF {
		//解析各种关键字，并放置到Statements中 如let,return
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()

	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

// 解析let声明,返回一个ast树
func (p *Parser) parseLetStatement() *ast.LetStatement {

	stmt := &ast.LetStatement{Token: p.curToken}

	//判断下一个是不是关键字
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	stmt.Name = &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}
	//判断下一个字符是不是等于（判断是不是赋值）
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}
	//判断最后是不是以分号结尾
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseReturnStatement() *ast.LetStatement {

	stmt := &ast.LetStatement{Token: p.curToken}
	//判断最后是不是以分号结尾
	//todo 这里如果结尾不是分号，那就会一直循环，后续应该要做处理
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()

	}
	return stmt
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {

	stmt := &ast.ExpressionStatement{Token: p.curToken}
	stmt.Expression = p.parseExpression()
	//判断最后是不是以分号结尾
	//todo 这里如果结尾不是分号，那就会一直循环，后续应该要做处理
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()

	}
	return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) Errors() []string {
	//返回errors，目标应该是不让用户直接修改这个值
	return p.errors
}

// 追加异常
func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s,got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)

}

func (p *Parser) registerPrefix(t token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[t] = fn

}

func (p *Parser) registerInfix(t token.TokenType, fn infixParseFn) {
	p.infixParseFns[t] = fn

}
