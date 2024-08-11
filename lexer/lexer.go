package lexer

import (
	"fmt"
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
	l.position = l.readPosition // 将指针指向下一个字符
	l.readPosition += 1
	if l.readPosition > len(l.input) {
		l.ch = 0 // 将当前的ascii吗归为0，代表没找到这个字符
	} else {
		l.ch = l.input[l.position] // go里面直接用索引拿字符串的值，会返回一个ascii码
	}

}

func (l *Lexer) readNumber() string {
	//这里读取的是字节码
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func (l *Lexer) NextToken() token.Token {
	// 根据当前的标识符返回token对象
	var tok token.Token

	// 校验字符是否为空
	l.skipWhitespace()

	//校验当前的字符是不是下列符号中的一种，匹配的是ASCII码
	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			fmt.Println(string(l.ch))

			tok.Type = token.EQ
			tok.Literal = string(l.ch) + "="
			l.readChar()
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '!':
		if l.peekChar() == '=' {
			tok.Type = token.NOT_EQ
			tok.Literal = string(l.ch) + string(l.peekChar())
			l.readChar()
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
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
		tok.Type = token.EOF
		tok.Literal = ""
	default:
		//如果是首字母、下划线开头
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier() // 读取关键字，并赋值给tok
			tok.Type = token.LookupIdent(tok.Literal)
			return tok

		} else if isDigit(l.ch) {
			//如果是数字开头，将数字读取出来，写到token中
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	// 语法分析器读取下一个字符串
	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// 读取关键字：这里会返回字符串
func (l *Lexer) readIdentifier() string {

	position := l.position //记录当前的索引位置

	// 一直读取字符串直到下一个字符串不在关键字范围内
	for isLetter(l.ch) {
		//第一次索引是0
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
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
