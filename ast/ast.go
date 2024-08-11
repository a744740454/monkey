package ast

import "monky/token"

type Node interface {
	// TokenLiteral 该方法必须实现，用于返回词法单元的值
	TokenLiteral() string
}

// Statement 声明
type Statement interface {
	Node
	statementNode()
}

// Expression 表达式
type Expression interface {
	Node
	expressionNode()
}

// Program 当前程序的结构体，存放关键式声明
type Program struct {
	Statements []Statement
}

// LetStatement let的声明
type LetStatement struct {
	Token token.Token //词法单元，包含类型以及值
	Name  *Identifier // 绑定的标识符
	value Expression  //值的表达式
}

// Identifier 标识符
type Identifier struct {
	Token token.Token
	Value string
}

func (p Program) statementNode() {
	return
}

func (p Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (i *Identifier) expressionNode() {

}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (ls *LetStatement) statementNode() {
	return
}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}
