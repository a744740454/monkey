package lexer

import (
	"monky/token"
)

type Lexer struct {
	input        string // 用户的输入
	position     int    // 当前的字符位置
	readPosition int    // 当前的索引，指向chr对应的后一个索引
	ch           byte   // 当前的字符
}

// New 初始化一个词法分析器对象
func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	//这里读取的是字节码
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.position] // go里面直接用索引拿字符串的值，会返回一个ascii码
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l Lexer) NextToken() token.Token {
	// 根据当前的标识符返回token对象
	var tok token.Token

	//校验当前的字符是不是下列符号中的一种，匹配的是ASCII码
	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = token.EOF
		tok.Literal = ""
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier() // 读取关键字，并赋值给tok
			tok.Type = token.LookupIdent(tok.Literal)
			return tok

		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	// 语法分析器读取下一个字符串
	l.readChar()
	return tok
}

// 读取关键字：这里会返回字符串
func (l Lexer) readIdentifier() string {

	position := l.position //记录当前的索引位置

	// 一直读取字符串直到下一个字符串不在关键字范围内
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// 用于判断输入的字符是不是命名规则中的一员
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type: tokenType, Literal: string(ch),
	}
}
